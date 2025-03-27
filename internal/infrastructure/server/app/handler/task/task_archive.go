package task

import (
	"encoding/json"
	"net/http"

	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/apperror"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/request/task"
)

type ArchiveTaskHandler struct {
	appHandler handler.AppHandlerInterface
}

func NewArchiveTaskHandler(appHandler handler.AppHandlerInterface) *ArchiveTaskHandler {
	return &ArchiveTaskHandler{appHandler}
}

func (h *ArchiveTaskHandler) GetPattern() string {
	return "POST /task/archive"
}

func (h *ArchiveTaskHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var requestParams task.ArchiveTaskRequestParams
	err := json.NewDecoder(r.Body).Decode(&requestParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}
	err = h.appHandler.Validator().Validate(&requestParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	archiveParams := db.TaskArchiveParams{
		ID:       requestParams.ID,
		Archived: requestParams.Archived,
	}
	err = h.appHandler.GetQuerier().TaskArchive(h.appHandler.Context(), archiveParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(archiveParams.ID)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	return nil
}
