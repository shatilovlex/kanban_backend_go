// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	Board(ctx context.Context, projectID pgtype.UUID) ([]*BoardRow, error)
	ListAdd(ctx context.Context, arg ListAddParams) error
	ListRemove(ctx context.Context, id pgtype.UUID) error
	ProjectArchive(ctx context.Context, arg ProjectArchiveParams) error
	ProjectCreate(ctx context.Context, arg ProjectCreateParams) error
	ProjectList(ctx context.Context) ([]*ProjectListRow, error)
	RenameList(ctx context.Context, arg RenameListParams) error
	SaveListOrder(ctx context.Context, arg SaveListOrderParams) error
	TaskCreate(ctx context.Context, arg TaskCreateParams) error
}

var _ Querier = (*Queries)(nil)
