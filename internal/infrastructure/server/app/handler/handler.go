package handler

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/appvalidator"
)

type Handler struct {
	ctx       context.Context
	connect   *pgxpool.Pool
	querier   db.Querier
	validator *appvalidator.AppValidator
}

func NewMyHandler(ctx context.Context, connect *pgxpool.Pool) *Handler {
	querier := db.New(connect)
	validator := appvalidator.NewAppValidator()
	return &Handler{ctx: ctx, connect: connect, querier: querier, validator: validator}
}

func (h *Handler) GetQuerier() db.Querier {
	return h.querier
}

func (h *Handler) Connect() *pgxpool.Pool {
	return h.connect
}

func (h *Handler) Context() context.Context {
	return h.ctx
}

func (h *Handler) Validator() *appvalidator.AppValidator {
	return h.validator
}
