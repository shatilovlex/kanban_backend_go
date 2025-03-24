package list

import (
	"encoding/json"
	"net/http"

	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/apperror"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/request/list"
)

type RenameListHandler struct {
	appHandler handler.AppHandlerInterface
}

func NewRenameListHandler(appHandler handler.AppHandlerInterface) *RenameListHandler {
	return &RenameListHandler{appHandler}
}

func (h *RenameListHandler) GetPattern() string {
	return "POST /v1/renameList"
}

func (h *RenameListHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	renameListRequestParams := list.RenameListRequestParams{}
	err := json.NewDecoder(r.Body).Decode(&renameListRequestParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	err = h.appHandler.Validator().Validate(&renameListRequestParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	renameListParams := db.RenameListParams{
		ID:   renameListRequestParams.ID,
		Name: &renameListRequestParams.Name,
	}

	err = h.appHandler.GetQuerier().RenameList(h.appHandler.Context(), renameListParams)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}
