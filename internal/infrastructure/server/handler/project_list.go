package handler

import (
	"encoding/json"
	"net/http"

	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/myHandler"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/statusError"
)

type ProjectListHandler struct {
	appHandler *myHandler.MyHandler
}

func NewProjectListHandler(appHandler *myHandler.MyHandler) *ProjectListHandler {
	return &ProjectListHandler{appHandler}
}

func (h *ProjectListHandler) GetPattern() string {
	return "GET /v1/projects"
}

func (h *ProjectListHandler) Handle(w http.ResponseWriter, _ *http.Request) error {
	res, err := h.appHandler.GetQuerier().ProjectList(h.appHandler.Context())
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	return nil
}
