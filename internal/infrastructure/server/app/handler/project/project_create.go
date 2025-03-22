package project

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/apperror"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/request/project"
)

type CreateProjectHandler struct {
	appHandler *handler.MyHandler
}

func NewCreateProjectHandler(appHandler *handler.MyHandler) *CreateProjectHandler {
	return &CreateProjectHandler{appHandler}
}

func (h *CreateProjectHandler) GetPattern() string {
	return "POST /project/create"
}

func (h *CreateProjectHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var id pgtype.UUID
	projectRequestParams := project.CreateProjectRequestParams{}
	err := json.NewDecoder(r.Body).Decode(&projectRequestParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	err = h.appHandler.Validator().Validate(&projectRequestParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	err = id.Scan(uuid.New().String())
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	err = h.appHandler.GetQuerier().ProjectCreate(h.appHandler.Context(), db.ProjectCreateParams{
		ID:          id,
		Name:        &projectRequestParams.Name,
		Description: &projectRequestParams.Description,
	})
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(id)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	return nil
}
