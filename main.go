package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pedrompeixoto/person-budget-api/internal/budget"
	"github.com/pedrompeixoto/person-budget-api/store"
)

var validCommands = map[string]bool{
	"start":   true,
	"migrate": true,
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), `Usage:
  go run . <command> [subcommand] [options]

Commands:
  start                 Start the HTTP server
  migrate create <name> Create a new migration file
  migrate run           Apply pending migrationsmigrate   Run database migrations
`)
	}

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	command := flag.Arg(0)
	if !validCommands[command] {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		flag.Usage()
		os.Exit(1)
	}

	pbStore, err := store.New("personal_budget.db")
	if err != nil {
		log.Fatal(err)
	}

	switch command {
	case "start":
		startServer(pbStore)
	case "migrate":
		handleMigrate(flag.Args()[1:], pbStore)
	}

}

func handleMigrate(args []string, pbStore *store.Store) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: go run . migrate <create|run> [name]")
		os.Exit(1)
	}

	switch args[0] {
	case "create":
		if len(args) < 2 {
			log.Fatal("please provide a migration name, e.g. `go run . migrate create add_users_table`")
		}
		store.CreateMigration(args[1])
	case "run":
		store.Migrate(pbStore)
	default:
		log.Fatalf("unknown migrate subcommand: %s", args[0])
	}
}

func startServer(store *store.Store) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "openapi.yaml")
	})
	fs := http.FileServer(http.Dir("./swagger-ui"))
	http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", fs))

	// mount the budget handler
	budgetHandler := budget.NewBudgetHandler(store)
	http.Handle("/budget", budgetHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
