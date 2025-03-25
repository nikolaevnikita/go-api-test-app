package repository

type Repository[T any] interface {
	Get(id string) (*T, error)
	GetAll() ([]*T, error)
	Create(item T) error
	Update(id string, item T) error
	Delete(id string) error
}
  