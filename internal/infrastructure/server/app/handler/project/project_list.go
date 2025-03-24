package project

import (
	"encoding/json"
	"net/http"

	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/apperror"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
)

type GetProjectListHandler struct {
	appHandler handler.AppHandlerInterface
}

func NewProjectListHandler(appHandler handler.AppHandlerInterface) *GetProjectListHandler {
	return &GetProjectListHandler{appHandler}
}

func (h *GetProjectListHandler) GetPattern() string {
	return "GET /v1/projects"
}

func (h *GetProjectListHandler) Handle(w http.ResponseWriter, _ *http.Request) error {
	res, err := h.appHandler.GetQuerier().ProjectList(h.appHandler.Context())
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	return nil
}
