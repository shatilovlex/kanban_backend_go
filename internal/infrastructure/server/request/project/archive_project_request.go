package project

import "github.com/jackc/pgx/v5/pgtype"

type ArchiveProjectRequestParams struct {
	ID       pgtype.UUID `json:"id" validate:"required,uuid4"`
	Archived bool        `json:"archived" validate:"required"`
}
