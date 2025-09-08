package budget

import (
	"net/http"
	"github.com/pedrompeixoto/person-budget-api/store"
)

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

func (h *BudgetHandler) create(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("POST budget"))
}

func (h *BudgetHandler) update(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("PUT budget"))
}

func (h *BudgetHandler) delete(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("DELETE budget"))
}

