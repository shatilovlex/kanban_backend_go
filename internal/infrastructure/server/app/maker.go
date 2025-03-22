package app

import (
	"net/http"

	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/myHandler"
)

type MuxHandlerInterface interface {
	GetPattern() string
	myHandler.MyHandlerInterface
}

type MakerAppMux struct {
	handlers []MuxHandlerInterface
}

func NewMakerAppMux(handlers []MuxHandlerInterface) *MakerAppMux {
	return &MakerAppMux{handlers: handlers}
}

func (m *MakerAppMux) MakeHandlers(mux *http.ServeMux) {
	for _, handler := range m.handlers {
		mux.Handle(handler.GetPattern(), m.CustomHandler(handler.Handle))
	}
}

func (m *MakerAppMux) CustomHandler(f func(http.ResponseWriter, *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
