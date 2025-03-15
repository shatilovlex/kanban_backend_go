package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
)

type BoardHandler struct {
	ctx     context.Context
	querier db.Querier
}

func NewBoardHandler(ctx context.Context, connect *pgxpool.Pool) *BoardHandler {
	querier := db.New(connect)
	return &BoardHandler{ctx: ctx, querier: querier}
}

func (h *BoardHandler) GetQuerier() db.Querier {
	return h.querier
}

func (h *BoardHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var res []*db.BoardRow
	projectID := pgtype.UUID{}
	err := projectID.Scan(r.URL.Query().Get("project_id"))
	if err != nil {
		http.Error(w, "Invalid project_id", http.StatusBadRequest)
		log.Printf("Error get project_id: %v", err)
		return
	}
	res, err = h.GetQuerier().Board(h.ctx, projectID)
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
