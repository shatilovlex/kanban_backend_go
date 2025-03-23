package handler

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/appvalidator"
)

type AppHandlerMock struct {
	ctx     context.Context
	connect *pgxpool.Pool
	querier db.Querier
}

func NewHandlerMock(ctx context.Context, querier db.Querier) *AppHandlerMock {
	return &AppHandlerMock{ctx: ctx, connect: nil, querier: querier}
}

func (h *AppHandlerMock) GetQuerier() db.Querier {
	return h.querier
}

func (h *AppHandlerMock) Connect() *pgxpool.Pool {
	return h.connect
}

func (h *AppHandlerMock) Context() context.Context {
	return h.ctx
}

func (h *AppHandlerMock) Validator() *appvalidator.AppValidator {
	return appvalidator.NewAppValidator()
}
