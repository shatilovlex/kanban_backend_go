package handler

import (
	"encoding/json"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/statusError"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/myHandler"
)

type BoardHandler struct {
	appHandler *myHandler.MyHandler
}

func NewBoardHandler(appHandler *myHandler.MyHandler) *BoardHandler {
	return &BoardHandler{appHandler}
}

func (h *BoardHandler) GetPattern() string {
	return "GET /v1/board"
}

func (h *BoardHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var res []*db.BoardRow
	projectID := pgtype.UUID{}
	err := projectID.Scan(r.URL.Query().Get("project_id"))
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusBadRequest)
	}
	res, err = h.appHandler.GetQuerier().Board(h.appHandler.Context(), projectID)
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusBadRequest)
	}
	return nil
}
