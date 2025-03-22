package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/myHandler"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/statusError"
)

type ProjectRequestParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateProjectHandler struct {
	appHandler *myHandler.MyHandler
}

func NewCreateProjectHandler(appHandler *myHandler.MyHandler) *CreateProjectHandler {
	return &CreateProjectHandler{appHandler}
}

func (h *CreateProjectHandler) GetPattern() string {
	return "POST /project/create"
}

func (h *CreateProjectHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var id pgtype.UUID
	projectRequestParams := ProjectRequestParams{}
	err := json.NewDecoder(r.Body).Decode(&projectRequestParams)
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusBadRequest)
	}

	if !isValidParams(&projectRequestParams) {
		return statusError.WithHTTPStatus(err, http.StatusBadRequest)
	}

	err = id.Scan(uuid.New().String())
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	err = h.appHandler.GetQuerier().ProjectCreate(h.appHandler.Context(), db.ProjectCreateParams{
		ID:          id,
		Name:        &projectRequestParams.Name,
		Description: &projectRequestParams.Description,
	})
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(id)
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	return nil
}

func isValidParams(params *ProjectRequestParams) bool {
	if params.Name == "" {
		return false
	}

	if params.Description == "" {
		return false
	}

	return true
}
