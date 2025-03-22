package myHandler

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
)

type MyHandlerInterface interface {
	Handle(w http.ResponseWriter, r *http.Request) error
}

type MyHandler struct {
	ctx     context.Context
	connect *pgxpool.Pool
	querier db.Querier
}

func NewMyHandler(ctx context.Context, connect *pgxpool.Pool) *MyHandler {
	querier := db.New(connect)
	return &MyHandler{ctx: ctx, connect: connect, querier: querier}
}

func (h *MyHandler) GetQuerier() db.Querier {
	return h.querier
}

func (h *MyHandler) Connect() *pgxpool.Pool {
	return h.connect
}

func (h *MyHandler) Context() context.Context {
	return h.ctx
}
