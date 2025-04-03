package repository

import (
	"github.com/nikolaevnikita/go-api-test-app/internal/domain/models"
	"github.com/nikolaevnikita/go-api-test-app/internal/domain/errors"
)

type ItemID = string

type InMemoryRepository[T models.Identifiable] struct {
	storage map[ItemID]*T
}

// MARK: Fabric

func NewTaskInMemoryRepository() *InMemoryRepository[models.Task] {
	return &InMemoryRepository[models.Task]{
		storage: make(map[ItemID]*models.Task),
	}
}

func NewUserInMemoryRepository() *InMemoryRepository[models.User] {
	return &InMemoryRepository[models.User]{
		storage: make(map[ItemID]*models.User),
	}
}

// MARK: CRUD operations

func (r *InMemoryRepository[T]) Get(id ItemID) (*T, error) {
	if err := r.checkIsFound(id); err != nil {
		return nil, err
	}
	return r.storage[id], nil
}

func (r *InMemoryRepository[T]) GetAll() ([]*T, error) {
	var items []*T
	for _, item := range r.storage {
		items = append(items, item)
	}
	return items, nil   
}

func (r *InMemoryRepository[T]) Create(item T) error {
	id := item.ID()
	if err := r.checkIsNotFound(id); err != nil {
		return err
	}
	r.storage[id] = &item
	return nil
}

func (r *InMemoryRepository[T]) Update(id ItemID, item T) error {
	if err := r.checkIsFound(id); err != nil {
		return err
	}
	r.storage[id] = &item
	return nil
}

func (r *InMemoryRepository[T]) Delete(id ItemID) error {
	if err := r.checkIsFound(id); err != nil {
		return err
	}
	delete(r.storage, id)
	return nil
}

// MARK: Private methods

func (r *InMemoryRepository[T]) checkIsFound(id ItemID) error {
	_, ok := r.storage[id]
	if ok {
		return nil
	}
	return errors.ErrNotFound
}

func (r *InMemoryRepository[T]) checkIsNotFound(id ItemID) error {
	_, ok := r.storage[id]
	if ok {
		return errors.ErrAlreadyExists
	}
	return nil
}
