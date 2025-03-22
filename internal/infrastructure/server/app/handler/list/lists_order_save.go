package list

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/apperror"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
)

type SaveOrderHandler struct {
	appHandler *handler.Handler
}

func NewSaveOrderHandler(appHandler *handler.Handler) *SaveOrderHandler {
	return &SaveOrderHandler{appHandler}
}

func (h *SaveOrderHandler) GetPattern() string {
	return "POST /v1/saveListOrder"
}

func (h *SaveOrderHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	tx, err := h.appHandler.Connect().BeginTx(h.appHandler.Context(), pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	defer tx.Rollback(h.appHandler.Context())

	var params []db.SaveListOrderParams
	var projectID pgtype.UUID
	var board []*db.BoardRow
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}
	for _, param := range params {
		err = h.appHandler.GetQuerier().SaveListOrder(h.appHandler.Context(), param)
		projectID = param.ProjectID
		if err != nil {
			return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
		}
	}

	err = tx.Commit(h.appHandler.Context())
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	board, err = h.appHandler.GetQuerier().Board(h.appHandler.Context(), projectID)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(board)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	return nil
}
