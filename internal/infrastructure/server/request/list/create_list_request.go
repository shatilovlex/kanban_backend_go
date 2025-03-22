package list

import "github.com/jackc/pgx/v5/pgtype"

type CreateListRequestParams struct {
	Name      string      `json:"name" validate:"required"`
	ProjectID pgtype.UUID `json:"projectI" validate:"required"`
	Sort      int32       `json:"sort" validate:"required"`
}
