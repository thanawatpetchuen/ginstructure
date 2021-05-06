package app

import (
	"context"
	"log"
	"net/http"

	"github.com/thanawatpetchuen/ginstructure/handler"
	"github.com/thanawatpetchuen/ginstructure/internal/config"
	"github.com/thanawatpetchuen/ginstructure/internal/server"
	"github.com/thanawatpetchuen/ginstructure/model"
	"github.com/thanawatpetchuen/ginstructure/router"
	"github.com/thanawatpetchuen/ginstructure/service"
)

type App struct {
	router *http.Server
	cfg    *model.Config
}

func NewApp() *App {
	cfg := model.Config{}

	config.LoadConfig(&cfg, "config")

	todoService := service.NewTodoService()
	handlers := handler.NewHandler(todoService)
	httpRouter := router.NewHTTPRouter(handlers)
	httpServer := server.NewServer(cfg.Http.Port, httpRouter)

	return &App{
		router: httpServer,
		cfg:    &cfg,
	}
}

func (app *App) Start() error {
	log.Printf("Start service: %s (%s) \n", app.cfg.ServiceName, app.cfg.Env)
	return app.router.ListenAndServe()
}

func (app *App) Shutdown(ctx context.Context) error {
	return app.router.Shutdown(ctx)
}
