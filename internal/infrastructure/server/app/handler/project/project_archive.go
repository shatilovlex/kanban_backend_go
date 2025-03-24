package project

import (
	"encoding/json"
	"net/http"

	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/apperror"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/request/project"
)

type ArchiveProjectHandler struct {
	appHandler handler.AppHandlerInterface
}

func NewArchiveProjectHandler(appHandler handler.AppHandlerInterface) *ArchiveProjectHandler {
	return &ArchiveProjectHandler{appHandler}
}

func (h *ArchiveProjectHandler) GetPattern() string {
	return "POST /project/archive"
}

func (h *ArchiveProjectHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var requestParams project.ArchiveProjectRequestParams
	err := json.NewDecoder(r.Body).Decode(&requestParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}
	err = h.appHandler.Validator().Validate(&requestParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	projectRequestParams := db.ProjectArchiveParams{
		ID:       requestParams.ID,
		Archived: requestParams.Archived,
	}
	err = h.appHandler.GetQuerier().ProjectArchive(h.appHandler.Context(), projectRequestParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(projectRequestParams.ID)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	return nil
}
