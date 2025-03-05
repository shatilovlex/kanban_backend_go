package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
)

type ArchiveProjectHandler struct {
	ctx     context.Context
	querier db.Querier
}

func NewArchiveProjectHandler(ctx context.Context, connect *pgxpool.Pool) *ArchiveProjectHandler {
	querier := db.New(connect)
	return &ArchiveProjectHandler{ctx: ctx, querier: querier}
}

func (h *ArchiveProjectHandler) GetQuerier() db.Querier {
	return h.querier
}

func (h *ArchiveProjectHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var projectRequestParams db.ProjectArchiveParams
	err := json.NewDecoder(r.Body).Decode(&projectRequestParams)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.GetQuerier().ProjectArchive(h.ctx, projectRequestParams)
	if err != nil {
		http.Error(w, "Failed change status project", http.StatusInternalServerError)
		log.Printf("Error change status project: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(projectRequestParams.ID)
	log.Println(projectRequestParams)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
