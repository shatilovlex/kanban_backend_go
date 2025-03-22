package project

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/apperror"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
)

type ArchiveProjectHandler struct {
	appHandler *handler.Handler
}

func NewArchiveProjectHandler(appHandler *handler.Handler) *ArchiveProjectHandler {
	return &ArchiveProjectHandler{appHandler}
}

func (h *ArchiveProjectHandler) GetPattern() string {
	return "POST /project/archive"
}

func (h *ArchiveProjectHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var projectRequestParams db.ProjectArchiveParams
	err := json.NewDecoder(r.Body).Decode(&projectRequestParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	err = h.appHandler.GetQuerier().ProjectArchive(h.appHandler.Context(), projectRequestParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(projectRequestParams.ID)
	log.Println(projectRequestParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	return nil
}
