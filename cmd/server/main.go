package main

import (
	"context"
	"log"

	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app"
)

func main() {
	ctx := context.Background()
	mainApp, err := app.NewApp(ctx)
	if err != nil {
		log.Panicln(err.Error())
	}
	mainApp.Start()
}
