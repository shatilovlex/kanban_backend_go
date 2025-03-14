package app

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/config"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/handler"
	"github.com/shatilovlex/kanban_backend_go/pkg/pgconnect"
)

type App struct {
	ctx    context.Context
	config *config.Cfg
	db     *pgxpool.Pool
}

func NewApp(ctx context.Context) (*App, error) {
	conf, err := config.Init()
	if err != nil {
		return nil, err
	}

	db, err := pgconnect.NewDB(ctx, conf.DB)
	if err != nil {
		return nil, err
	}

	return &App{
		ctx:    ctx,
		config: conf,
		db:     db,
	}, nil
}

func (a *App) Start() {
	ctx, stop := signal.NotifyContext(a.ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	ip := flag.String("ip", a.config.HTTP.Host, "IP address")
	port := flag.String("port", a.config.HTTP.Port, "Port number")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	mux.HandleFunc("POST /project/create", handler.NewCreateProjectHandler(a.ctx, a.db).Handle)
	mux.HandleFunc("POST /project/archive", handler.NewArchiveProjectHandler(a.ctx, a.db).Handle)
	mux.HandleFunc("GET /projects", handler.NewProjectListHandler(a.ctx, a.db).Handle)

	addr := fmt.Sprintf("%v:%v", *ip, *port)
	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 2 * time.Second,
	}
	go func() {
		log.Printf("start receiving at: %v", addr)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen and serve returned err: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("got interruption signal")
	ctxT, cancel := context.WithTimeout(a.ctx, 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctxT); err != nil {
		log.Printf("error while shutting down http server: %s", err)
	}
	log.Println("final")
}
