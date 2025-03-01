package pgconnect

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/config"
)

func NewDB(ctx context.Context, dbCfg config.DB) (*pgxpool.Pool, error) {
	connConfig, err := pgx.ParseConfig(
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s?TimeZone=Europe/Moscow",
			dbCfg.User,
			dbCfg.Password,
			net.JoinHostPort(dbCfg.Host, strconv.Itoa(dbCfg.Port)),
			dbCfg.Database,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create DSN for DB connection: %w", err)
	}
	dbc, err := pgxpool.New(ctx, connConfig.ConnString())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB : %w", err)
	}
	if err = dbc.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return dbc, nil
}
