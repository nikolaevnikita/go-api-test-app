package repository

type ItemID = string

type Repository[T any] interface {
	Get(id ItemID) (*T, error)
	GetAll() ([]*T, error)
	Create(item T) error
	Update(id ItemID, item T) error
	Delete(id ItemID) error
}
  