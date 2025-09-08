package main

import (
	"log"
	"net/http"

	"github.com/pedrompeixoto/person-budget-api/internal/budget"
	"github.com/pedrompeixoto/person-budget-api/store"
)

func main() {
	pbStore, err := store.New("personal_budget.db")
	if err != nil {
		log.Fatal(err)
	}
	defer pbStore.DB.Close()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "openapi.yaml")
	})
	fs := http.FileServer(http.Dir("./swagger-ui"))
	http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", fs))

	// mount the budget handler
	budgetHandler := budget.NewBudgetHandler(pbStore)
	http.Handle("/budget", budgetHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
