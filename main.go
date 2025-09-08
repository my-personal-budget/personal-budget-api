package main

import (
	"log"
	"net/http"
	"github.com/pedrompeixoto/person-budget-api/store"
)

func main() {
    s, err := store.New("personal_budget.db")
    if err != nil {
        log.Fatal(err)
    }
    defer s.DB.Close()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


