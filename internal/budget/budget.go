package budget

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pedrompeixoto/person-budget-api/store"
)

type Budget struct {
	ID         string    `db:"id" json:"id"`
	Month      time.Time `db:"month" json:"month"`             // first day of month
	CategoryID int       `db:"category_id" json:"category_id"` // FK to categories
	Budget     float64   `db:"budget" json:"budget"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}

type CreateBudgetBody struct {
	Month      time.Time `json:"month"`
	CategoryID int       `json:"categoryId"`
	Budget     float64   `json:"budget"`
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
	var body CreateBudgetBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.store.CreateBudget(body.Month, body.CategoryID, body.Budget)
	if err != nil {
		http.Error(w, "failed to create budget", http.StatusInternalServerError)
		return
	}

	b := Budget{
		ID:         id,
		Month:      body.Month,
		CategoryID: body.CategoryID,
		Budget:     body.Budget,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

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
