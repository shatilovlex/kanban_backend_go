package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
)

type ProjectRequestParams struct {
	Name        string
	Description string
}

type CreateProjectHandler struct {
	ctx     context.Context
	querier db.Querier
}

func NewCreateProjectHandler(ctx context.Context, connect *pgxpool.Pool) *CreateProjectHandler {
	querier := db.New(connect)
	return &CreateProjectHandler{ctx: ctx, querier: querier}
}

func (h *CreateProjectHandler) GetQuerier() db.Querier {
	return h.querier
}

func (h *CreateProjectHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var id pgtype.UUID
	projectRequestParams := ProjectRequestParams{}
	err := json.NewDecoder(r.Body).Decode(&projectRequestParams)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = id.Scan(uuid.New().String())
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		log.Printf("Error create user: %v", err)
		return
	}

	err = h.GetQuerier().ProjectCreate(h.ctx, db.ProjectCreateParams{
		ID:          id,
		Name:        &projectRequestParams.Name,
		Description: &projectRequestParams.Description,
	})
	if err != nil {
		http.Error(w, "Failed to create project", http.StatusInternalServerError)
		log.Printf("Error create project: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(id)
	if err != nil {
		log.Println("Error encoding user response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
