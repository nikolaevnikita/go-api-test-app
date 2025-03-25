package server

import (
	"github.com/nikolaevnikita/go-api-test-app/internal/domain/models"

	"net/http"
	"github.com/gin-gonic/gin"
)

func (s *ServerApi) createTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindBodyWithJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTask, err := s.taskService.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}

func (s *ServerApi) updateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindBodyWithJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tID := c.Param("id")
	updatedTask, err := s.taskService.UpdateTask(tID, task)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (s *ServerApi) getTask(c *gin.Context) {
	tID := c.Param("id")

	task, err := s.taskService.GetTask(tID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (s *ServerApi) getTasks(c *gin.Context) {
	tasks, err := s.taskService.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (s *ServerApi) deleteTask(c *gin.Context) {
	tID := c.Param("id")

	if err := s.taskService.DeleteTask(tID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}