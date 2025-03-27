package task

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/apperror"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/request/task"
)

type CreateTaskHandler struct {
	appHandler handler.AppHandlerInterface
}

func NewCreateTaskHandler(appHandler handler.AppHandlerInterface) *CreateTaskHandler {
	return &CreateTaskHandler{appHandler}
}

func (h *CreateTaskHandler) GetPattern() string {
	return "POST /tasks/create"
}

func (h *CreateTaskHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var id pgtype.UUID
	createTaskRequest := task.CreateTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(&createTaskRequest)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	err = h.appHandler.Validator().Validate(&createTaskRequest)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	err = id.Scan(uuid.New().String())
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	createParams := db.TaskCreateParams{
		ID:          id,
		ListID:      createTaskRequest.ListID,
		Title:       &createTaskRequest.Title,
		Description: &createTaskRequest.Description,
		Sort:        &createTaskRequest.Sort,
	}
	err = h.appHandler.GetQuerier().TaskCreate(h.appHandler.Context(), createParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	return nil
}
