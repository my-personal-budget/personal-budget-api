package budget

import (
	"encoding/json"
	"net/http"

	"github.com/pedrompeixoto/person-budget-api/store"
)

type Budget struct {
    ID       string    `json:"id"`
    Month    string `json:"month"`    // e.g., "2025-09"
    Category string `json:"category"` // e.g., "Groceries"
    Budget   int    `json:"budget"`   // amount in integer
}

type BudgetHandler struct {
	store *store.Store
}

func NewBudgetHandler(s *store.Store) *BudgetHandler {
	return &BudgetHandler{store: s}
}

func (h *BudgetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	case http.MethodPost:
		h.create(w, r)
	case http.MethodPut:
		h.update(w, r)
	case http.MethodDelete:
		h.delete(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *BudgetHandler) get(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("GET budgets"))
}

func (h *BudgetHandler) create(w http.ResponseWriter, r *http.Request) {
    var b Budget
    if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
        http.Error(w, "invalid request body", http.StatusBadRequest)
        return
    }

    id, err := h.store.CreateBudget(b.Month, b.Category, b.Budget)
    if err != nil {
        http.Error(w, "failed to create budget", http.StatusInternalServerError)
        return
    }

    b.ID = id

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(b)
}

func (h *BudgetHandler) update(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("PUT budget"))
}

func (h *BudgetHandler) delete(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("DELETE budget"))
}

