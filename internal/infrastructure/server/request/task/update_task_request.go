package task

import "github.com/jackc/pgx/v5/pgtype"

type UpdateTaskRequest struct {
	Title       string      `json:"title" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Sort        int32       `json:"sort" validate:"required"`
	ID          pgtype.UUID `json:"id" validate:"required,uuid4"`
	ListID      pgtype.UUID `json:"listId" validate:"required,uuid4"`
}
