package list

import "github.com/jackc/pgx/v5/pgtype"

type RenameListRequestParams struct {
	Name string      `json:"name" validate:"required"`
	ID   pgtype.UUID `json:"id" validate:"required,uuid4"`
}
