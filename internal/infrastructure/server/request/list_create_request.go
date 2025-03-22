package request

import "github.com/jackc/pgx/v5/pgtype"

type ListRequestParams struct {
	Name      string      `json:"name" validate:"required"`
	ProjectID pgtype.UUID `json:"project_id" validate:"required"`
	Sort      int32       `json:"sort" validate:"required"`
}
