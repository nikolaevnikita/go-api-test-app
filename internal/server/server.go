package server

import (
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
}

func New(cfg config.Config, taskService *services.TaskService) *ServerApi {
	server := http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	router := gin.Default()
	server.Handler = router

	serverApi := ServerApi{
		server: &server,
		router: router,
		taskService: taskService,
	}

	serverApi.setupRoutes()

	return &serverApi
}

func (s *ServerApi) Start() error {
	return s.server.ListenAndServe()
}

// MARK: Private methods

func (serverApi *ServerApi) setupRoutes() {
	router := serverApi.router
	task := router.Group("/task")
	{
		task.GET("/:id", serverApi.getTask)
		task.PUT("/:id", serverApi.updateTask)
		task.DELETE("/:id", serverApi.deleteTask)
	}
	router.GET("/tasks", serverApi.getTasks)
	router.POST("/task", serverApi.createTask)
}