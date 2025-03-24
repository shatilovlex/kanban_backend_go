package list

import "github.com/jackc/pgx/v5/pgtype"

type RemoveListRequestParams struct {
	ID pgtype.UUID `json:"id" validate:"required,uuid4"`
}
