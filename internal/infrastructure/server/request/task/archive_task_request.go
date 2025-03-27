package task

import "github.com/jackc/pgx/v5/pgtype"

type ArchiveTaskRequestParams struct {
	ID       pgtype.UUID `json:"id" validate:"required,uuid4"`
	Archived bool        `json:"archived" validate:"required"`
}
