package task

import (
	"encoding/json"
	"net/http"

	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/apperror"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/request/task"
)

type UpdateTaskHandler struct {
	appHandler handler.AppHandlerInterface
}

func NewUpdateTaskHandler(appHandler handler.AppHandlerInterface) *UpdateTaskHandler {
	return &UpdateTaskHandler{appHandler}
}

func (h *UpdateTaskHandler) GetPattern() string {
	return "POST /tasks/update"
}

func (h *UpdateTaskHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	updateTaskRequest := task.UpdateTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(&updateTaskRequest)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	err = h.appHandler.Validator().Validate(&updateTaskRequest)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	updateParams := db.TaskUpdateParams{
		ID:          updateTaskRequest.ID,
		ListID:      updateTaskRequest.ListID,
		Title:       &updateTaskRequest.Title,
		Description: &updateTaskRequest.Description,
		Sort:        &updateTaskRequest.Sort,
	}

	err = h.appHandler.GetQuerier().TaskUpdate(h.appHandler.Context(), updateParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}
