package list

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/apperror"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/request/list"
)

type CreateListHandler struct {
	appHandler handler.AppHandlerInterface
}

func NewCreateListHandler(appHandler handler.AppHandlerInterface) *CreateListHandler {
	return &CreateListHandler{appHandler}
}

func (h *CreateListHandler) GetPattern() string {
	return "POST /v1/addList"
}

func (h *CreateListHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var id pgtype.UUID
	listRequestParams := list.CreateListRequestParams{}
	err := json.NewDecoder(r.Body).Decode(&listRequestParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}
	err = h.appHandler.Validator().Validate(&listRequestParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	err = id.Scan(uuid.New().String())
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	listAddParams := db.ListAddParams{
		ID:        id,
		ProjectID: listRequestParams.ProjectID,
		Name:      &listRequestParams.Name,
		Sort:      &listRequestParams.Sort,
	}
	err = h.appHandler.GetQuerier().ListAdd(h.appHandler.Context(), listAddParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return nil
}
