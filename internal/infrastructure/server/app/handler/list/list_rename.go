package list

import (
	"encoding/json"
	"net/http"

	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/apperror"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
)

type RenameListHandler struct {
	appHandler *handler.Handler
}

func NewRenameListHandler(appHandler *handler.Handler) *RenameListHandler {
	return &RenameListHandler{appHandler}
}

func (h *RenameListHandler) GetPattern() string {
	return "POST /v1/renameList"
}

func (h *RenameListHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	renameListParams := db.RenameListParams{}
	err := json.NewDecoder(r.Body).Decode(&renameListParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	if !h.isValidParams(&renameListParams) {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	err = h.appHandler.GetQuerier().RenameList(h.appHandler.Context(), renameListParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(renameListParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	return nil
}

func (h *RenameListHandler) isValidParams(_ *db.RenameListParams) bool {
	return true
}
