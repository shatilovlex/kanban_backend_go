// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"
)

type Querier interface {
	ProjectArchive(ctx context.Context, arg ProjectArchiveParams) error
	ProjectCreate(ctx context.Context, arg ProjectCreateParams) error
}

var _ Querier = (*Queries)(nil)
