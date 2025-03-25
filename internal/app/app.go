package app

import (
	"github.com/nikolaevnikita/go-api-test-app/internal/config"
	"github.com/nikolaevnikita/go-api-test-app/internal/repository"
	"github.com/nikolaevnikita/go-api-test-app/internal/server"
	"github.com/nikolaevnikita/go-api-test-app/internal/services"
)

type App struct {
	serverApi *server.ServerApi
}

func NewApp() *App {
	taskRepository := repository.NewTaskInMemoryRepository()
	taskService := services.NewTaskService(taskRepository)

	config := config.ReadConfig()

	serverApi := server.New(config, taskService)

	return &App{
		serverApi: serverApi,
	}
}

func (app *App) Start() error {
	if err := app.serverApi.Start(); err != nil {
		return err
	}
	return nil
}
