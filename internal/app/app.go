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
	config := config.ReadConfig()
	
	setupLogger(config)

	taskRepository, userRepository := provideRepositories(config)

	taskService := services.NewTaskService(taskRepository)
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

// MARK: Private methods

func setupLogger(config *config.Config) {
	log := logger.Get(config.Debug)
	log.Debug().Msg("logger was initialized")
	log.Debug().Str("host", config.Host).Int("port", config.Port).Send()
}

func provideRepositories(config *config.Config) (repository.Repository[models.Task], repository.Repository[models.User]) {
	log := logger.Get()
	
	var taskRepository repository.Repository[models.Task]
	var userRepository repository.Repository[models.User]

	// Try to use DB repositories
	if err := repository.Migrate(config.DbDsn, config.MigratePath); err != nil {
		log.Warn().Err(err).Msg("failed to migrate db")
	} else {
		dbTaskRepo, err := repository.NewPostgreSQLTaskRepository(context.Background(), config.DbDsn)
		if err != nil {
			log.Warn().Err(err).Msg("failed to connect to task db")
		} else {
			taskRepository = dbTaskRepo
		}

		dbUserRepo, err := repository.NewPostgreSQLUserRepository(context.Background(), config.DbDsn)
		if err != nil {
			log.Warn().Err(err).Msg("failed to connect to user db")
		} else {
			userRepository = dbUserRepo
		}
	}

	// If DB usage failed - use InMemory repositories
	if taskRepository == nil {
		log.Warn().Msg("use in-memory task storage")
		taskRepository = repository.NewTaskInMemoryRepository()
	}
	if userRepository == nil {
		log.Warn().Msg("use in-memory user storage")
		userRepository = repository.NewUserInMemoryRepository()
	}

	return taskRepository, userRepository
}
