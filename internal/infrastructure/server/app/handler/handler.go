package handler

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/appvalidator"
)

type AppHandlerInterface interface {
	GetQuerier() db.Querier
	Connect() *pgxpool.Pool
	Context() context.Context
	Validator() *appvalidator.AppValidator
}

type Handler struct {
	ctx     context.Context
	connect *pgxpool.Pool
	querier db.Querier
}

func NewAppHandler(ctx context.Context, connect *pgxpool.Pool) *Handler {
	querier := db.New(connect)
	return &Handler{ctx: ctx, connect: connect, querier: querier}
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
	return appvalidator.NewAppValidator()
}
