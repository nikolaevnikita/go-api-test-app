package server

import (
	"context"

	"github.com/nikolaevnikita/go-api-test-app/internal/config"
	"github.com/nikolaevnikita/go-api-test-app/internal/services"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerApi struct {
	server *http.Server
	router *gin.Engine
	taskService *services.TaskService
	userService *services.UserService
}

func New(
	cfg config.Config, 
	taskService *services.TaskService, 
	userService *services.UserService,
) *ServerApi {
	server := http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	router := gin.Default()
	server.Handler = router

	serverApi := ServerApi{
		server: &server,
		router: router,
		taskService: taskService,
		userService: userService,
	}

	serverApi.setupRoutes()

	return &serverApi
}

func (s *ServerApi) Start() error {
	return s.server.ListenAndServe()
}

func (s *ServerApi) Stop(ctx context.Context) []error {
	var stopErrors []error
	
	if err := s.server.Shutdown(ctx); err != nil {
        stopErrors = append(stopErrors, err)
    }

	if err := s.taskService.Stop(ctx); err != nil {
        stopErrors = append(stopErrors, err)
    }

	if err := s.userService.Stop(ctx); err != nil {
        stopErrors = append(stopErrors, err)
    }

	return stopErrors
}

// MARK: Private methods

func (serverApi *ServerApi) setupRoutes() {
	router := serverApi.router

	router.Use(GzipDecompressMiddleware())
	router.Use(GzipCompressMiddleware())

	// task routes
	task := router.Group("/task")
	{
		task.GET("/:id", serverApi.getTask)
		task.PUT("/:id", serverApi.updateTask)
		task.DELETE("/:id", serverApi.deleteTask)
	}
	router.GET("/tasks", serverApi.getTasks)
	router.POST("/task", serverApi.createTask)

	// user routes
	user := router.Group("/user")
	{
		user.GET("/:id", serverApi.getUser)
		user.PUT("/:id", serverApi.updateUserName)
		user.DELETE("/:id", serverApi.deleteUser)
	}
	router.GET("/users", serverApi.getUsers)
	router.POST("/user", serverApi.registerUser)

}