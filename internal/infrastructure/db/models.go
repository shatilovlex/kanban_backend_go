// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type KanbanProject struct {
	ID          pgtype.UUID `db:"id" json:"id"`
	Name        *string     `db:"name" json:"name"`
	Description *string     `db:"description" json:"description"`
	Archived    bool        `db:"archived" json:"archived"`
}
