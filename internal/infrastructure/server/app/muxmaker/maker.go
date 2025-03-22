package muxmaker

import (
	"log"
	"net/http"

	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/apperror"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
)

type MuxHandlerInterface interface {
	GetPattern() string
	handler.MyHandlerInterface
}

type MakerAppMux struct {
	handlers []MuxHandlerInterface
}

func NewMakerAppMux(handlers []MuxHandlerInterface) *MakerAppMux {
	return &MakerAppMux{handlers: handlers}
}

func (m *MakerAppMux) MakeHandlers(mux *http.ServeMux) {
	for _, appMuxHandler := range m.handlers {
		mux.Handle(appMuxHandler.GetPattern(), m.CustomHandler(appMuxHandler))
	}
}

func (m *MakerAppMux) CustomHandler(f handler.MyHandlerInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := f.Handle(w, r)
		if err != nil {
			log.Printf("Error code: %d, detail: '%s'", apperror.HTTPStatus(err), err.Error())
			http.Error(w, err.Error(), apperror.HTTPStatus(err))
		}
	})
}
