package list

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/apperror"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
)

type RemoveListRequestParams struct {
	ID pgtype.UUID `json:"id"`
}

type RemoveListHandler struct {
	appHandler *handler.MyHandler
}

func NewRemoveListHandler(appHandler *handler.MyHandler) *RemoveListHandler {
	return &RemoveListHandler{appHandler}
}

func (h *RemoveListHandler) GetPattern() string {
	return "POST /v1/removeList"
}

func (h *RemoveListHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	params := RemoveListRequestParams{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	err = h.appHandler.GetQuerier().ListRemove(h.appHandler.Context(), params.ID)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(params.ID)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	return nil
}
