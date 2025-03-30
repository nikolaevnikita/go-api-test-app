package server

import (
	"github.com/nikolaevnikita/go-api-test-app/internal/domain/models"

	"net/http"
	"github.com/gin-gonic/gin"
)

func (s *ServerApi) registerUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	registeredUser, err := s.userService.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, registeredUser.ToResponse())
}

func (s *ServerApi) getUser(c *gin.Context) {
	uID := c.Param("id")

	user, err := s.userService.GetUser(uID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user.ToResponse())
}

func (s *ServerApi) getUsers(c *gin.Context) {
	users, err := s.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var userResponses []models.UserResponse
	for _, user := range users {
		userResponse := user.ToResponse()
		userResponses = append(userResponses, userResponse)
	}

	c.JSON(http.StatusOK, userResponses)
}

func (s *ServerApi) updateUserName(c *gin.Context) {
	uID := c.Param("id")
	newName := c.Query("name")

	updatedUser, err := s.userService.UpdateUserName(uID, newName)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, updatedUser.ToResponse())
}

func (s *ServerApi) deleteUser(c *gin.Context) {
	uID := c.Param("id")

	if err := s.userService.DeleteUser(uID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
