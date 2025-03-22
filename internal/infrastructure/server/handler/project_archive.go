package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/myHandler"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/statusError"
)

type ArchiveProjectHandler struct {
	appHandler *myHandler.MyHandler
}

func NewArchiveProjectHandler(appHandler *myHandler.MyHandler) *ArchiveProjectHandler {
	return &ArchiveProjectHandler{appHandler}
}

func (h *ArchiveProjectHandler) GetPattern() string {
	return "POST /project/archive"
}

func (h *ArchiveProjectHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var projectRequestParams db.ProjectArchiveParams
	err := json.NewDecoder(r.Body).Decode(&projectRequestParams)
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusBadRequest)
	}

	err = h.appHandler.GetQuerier().ProjectArchive(h.appHandler.Context(), projectRequestParams)
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(projectRequestParams.ID)
	log.Println(projectRequestParams)
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	return nil
}
