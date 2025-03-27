package board

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/apperror"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/response/board"
)

type GetBoardHandler struct {
	appHandler handler.AppHandlerInterface
}

func NewGetBoardHandler(appHandler handler.AppHandlerInterface) *GetBoardHandler {
	return &GetBoardHandler{appHandler}
}

func (h *GetBoardHandler) GetPattern() string {
	return "GET /v1/board"
}

func (h *GetBoardHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var (
		listsRows []*db.BoardListsRow
		tasksRows []*db.BoardTasksRow
	)
	projectID := pgtype.UUID{}
	err := projectID.Scan(r.URL.Query().Get("project_id"))
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}
	listsRows, err = h.appHandler.GetQuerier().BoardLists(h.appHandler.Context(), projectID)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}

	res := make([]board.GetBoardResponse, 0, len(listsRows))
	for _, listRow := range listsRows {
		tasksRows, err = h.appHandler.GetQuerier().BoardTasks(h.appHandler.Context(), listRow.ID)
		if err != nil {
			return apperror.WithHTTPStatus(err, http.StatusInternalServerError)
		}
		res = append(res, board.GetBoardResponse{
			ID:        listRow.ID,
			Name:      *listRow.Name,
			ProjectID: listRow.ProjectID,
			Sort:      *listRow.Sort,
			Tasks:     tasksRows,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return apperror.WithHTTPStatus(err, http.StatusBadRequest)
	}
	return nil
}
