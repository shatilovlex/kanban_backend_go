package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
)

type ProjectListHandler struct {
	ctx     context.Context
	querier db.Querier
}

func NewProjectListHandler(ctx context.Context, connect *pgxpool.Pool) *ProjectListHandler {
	querier := db.New(connect)
	return &ProjectListHandler{ctx: ctx, querier: querier}
}

func (h *ProjectListHandler) GetQuerier() db.Querier {
	return h.querier
}

func (h *ProjectListHandler) Handle(w http.ResponseWriter, _ *http.Request) {
	res, err := h.GetQuerier().ProjectList(h.ctx)
	if err != nil {
		http.Error(w, "Failed change status project", http.StatusInternalServerError)
		log.Printf("Error change status project: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
