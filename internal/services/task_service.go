package services

import (
	"context"

	"github.com/nikolaevnikita/go-api-test-app/internal/domain/errors"
	"github.com/nikolaevnikita/go-api-test-app/internal/domain/models"
	"github.com/nikolaevnikita/go-api-test-app/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TaskService struct {
	repository repository.Repository[models.Task]
	validator *validator.Validate
}

func NewTaskService(repository repository.Repository[models.Task]) *TaskService {
	validator := validator.New()
	return &TaskService{
		repository: repository,
		validator: validator,
	}
}

// MARK: Business Logic

func (ts *TaskService) CreateTask(task models.Task) (*models.Task, error) {
	// check title uniqueness
	storedTasks, err := ts.repository.GetAll()
	if err != nil {
		return nil, err
	}
	for _, storedTask := range storedTasks {
		if storedTask.Title == task.Title {
			return nil, errors.ErrAlreadyExists
		}
	}

	// store in repository
	tID := uuid.New().String()
	task.TID = tID
	if err := ts.repository.Create(task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (ts *TaskService) UpdateTask(tID string, task models.Task) (*models.Task, error) {
	task.TID = tID
	// TODO: Сделать апдейт отдельных полей, не затирая недостающие
	if err := ts.repository.Update(task.TID, task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (ts *TaskService) GetTask(tID string) (*models.Task, error) {
	task, err := ts.repository.Get(tID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (ts *TaskService) GetTasks() ([]*models.Task, error) {
	tasks, err := ts.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ts *TaskService) DeleteTask(tID string) error {
	if err := ts.repository.Delete(tID); err != nil {
		return err
	}
	return nil
}

// MARK: Stop

func (ts *TaskService) Stop(ctx context.Context) error {
	return ts.repository.Stop(ctx)
}