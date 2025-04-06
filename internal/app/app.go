package app

import (
	"github.com/nikolaevnikita/go-api-test-app/internal/config"
	"github.com/nikolaevnikita/go-api-test-app/internal/domain/models"
	"github.com/nikolaevnikita/go-api-test-app/internal/logger"
	"github.com/nikolaevnikita/go-api-test-app/internal/repository"
	"github.com/nikolaevnikita/go-api-test-app/internal/server"
	"github.com/nikolaevnikita/go-api-test-app/internal/services"

	"context"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

type App struct {
	serverApi *server.ServerApi
}

func NewApp() *App {
	config, err := config.ReadConfig()

	log := logger.Get(config.Debug)
	log.Debug().Msg("logger was initialized")
	log.Debug().Str("host", config.Host).Int("port", config.Port).Send()

	if err != nil {
		log.Warn().Err(err).Send()
	}

	var taskRepository repository.Repository[models.Task]

	// разве сначала не надо сделать миграцию?
	dbTaskRepo, err := repository.NewPostgreSQLTaskRepository(context.Background(), config.DbDsn)
	if err != nil {
		log.Warn().Err(err).Msg("failed to connect to db, use in-memory storage")
		taskRepository = repository.NewTaskInMemoryRepository()
	} else {
		if err := repository.Migrate(config.DbDsn, config.MigratePath); err != nil {
			log.Warn().Err(err).Msg("failed to migrate db, use in-memory storage")
			taskRepository = repository.NewTaskInMemoryRepository()
		} else {
			taskRepository = dbTaskRepo
		}
	}

	taskService := services.NewTaskService(taskRepository)

	userRepository := repository.NewUserInMemoryRepository()
	userService := services.NewUserService(userRepository)	

	serverApi := server.New(*config, taskService, userService)

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
