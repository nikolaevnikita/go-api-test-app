package services

import (
	"github.com/nikolaevnikita/go-api-test-app/internal/domain/errors"
	"github.com/nikolaevnikita/go-api-test-app/internal/domain/models"
	"github.com/nikolaevnikita/go-api-test-app/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type UserService struct {
	repository repository.Repository[models.User]
	validator *validator.Validate
}

func NewUserService(repository repository.Repository[models.User]) *UserService {
	validator := validator.New()
	return &UserService {
		repository: repository,
		validator: validator,
	}
}

// MARK: Business Logic

func (us *UserService) RegisterUser(user models.User) (*models.User, error) {
	// validate fields rules
	if err := us.validator.Struct(user); err != nil {
		return nil, err
	}

	// check email uniqueness
	storedUsers, err := us.repository.GetAll()
	if err != nil {
		return nil, err
	}
	for _, storedUser := range storedUsers {
		if storedUser.Email == user.Email {
			return nil, fmt.Errorf("the email has already been registered: %w", errors.ErrAlreadyExists)
		}
	}

	// convert password to hash
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashBytes)

	// store in repository with unique id
	uID := uuid.New().String()
	user.UID = uID
	if err := us.repository.Create(user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserService) GetUser(uID string) (*models.User, error) {
	user, err := us.repository.Get(uID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetUsers() ([]*models.User, error) {
	users, err := us.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) UpdateUserName(uID string, newName string) (*models.User, error) {
	user, err := us.repository.Get(uID)
	if err != nil {
		return nil, err
	}

	user.Name = newName
	us.repository.Update(uID, *user)
	return user, nil
}

func (us *UserService) DeleteUser(uID string) error {
	if err := us.repository.Delete(uID); err != nil {
		return err
	}
	return nil
}


