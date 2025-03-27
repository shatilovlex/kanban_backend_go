package board

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
)

type GetBoardResponse struct {
	Name      string              `json:"name"`
	Tasks     []*db.BoardTasksRow `json:"tasks"`
	Sort      int32               `json:"sort"`
	ProjectID pgtype.UUID         `json:"projectId"`
	ID        pgtype.UUID         `json:"id"`
}
