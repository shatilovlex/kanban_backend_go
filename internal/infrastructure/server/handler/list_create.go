package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/myHandler"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/statusError"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/request"
)

type CreateListHandler struct {
	appHandler *myHandler.MyHandler
}

func NewCreateListHandler(appHandler *myHandler.MyHandler) *CreateListHandler {
	return &CreateListHandler{appHandler}
}

func (h *CreateListHandler) GetPattern() string {
	return "POST /v1/addList"
}

func (h *CreateListHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	validate := validator.New()
	var id pgtype.UUID
	listRequestParams := request.ListRequestParams{}
	err := json.NewDecoder(r.Body).Decode(&listRequestParams)
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusBadRequest)

	}
	err = validate.StructCtx(h.appHandler.Context(), &listRequestParams)
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusBadRequest)
	}

	err = id.Scan(uuid.New().String())
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	listAddParams := db.ListAddParams{
		ID:        id,
		ProjectID: listRequestParams.ProjectID,
		Name:      &listRequestParams.Name,
		Sort:      &listRequestParams.Sort,
	}
	err = h.appHandler.GetQuerier().ListAdd(h.appHandler.Context(), listAddParams)
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(listAddParams)
	if err != nil {
		return statusError.WithHTTPStatus(err, http.StatusInternalServerError)
	}
	return nil
}
