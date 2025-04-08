package repository

import (
	"github.com/nikolaevnikita/go-api-test-app/internal/domain/models"

	"context"
	"time"
	"github.com/jackc/pgx/v5"
)

type PostgreSQLUserRepository struct {
	db *pgx.Conn
}

// MARK: Fabric

func NewPostgreSQLUserRepository(ctx context.Context, connString string) (*PostgreSQLUserRepository, error) {
	db, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	userRepository := PostgreSQLUserRepository {
		db: db,
	}

	return &userRepository, nil
}

// MARK: CRUD operations

func (r *PostgreSQLUserRepository) Get(id ItemID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	row := r.db.QueryRow(ctx, "SELECT * from users WHERE uid = $1", id)

	var user models.User
	err := row.Scan(&user.UID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgreSQLUserRepository) GetAll() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *PostgreSQLUserRepository) Create(item models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	insertUserRequest := "INSERT INTO users (uid, name, email, password) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(ctx, insertUserRequest, item.UID, item.Name, item.Email, item.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgreSQLUserRepository) Update(id ItemID, item models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	updateUserRequest := "UPDATE users SET name = $2, email = $3, password = $4 WHERE uid = $1"
	_, err := r.db.Exec(ctx, updateUserRequest, id, item.Name, item.Email, item.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgreSQLUserRepository) Delete(id ItemID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	deleteUserRequest := "DELETE FROM users WHERE uid = $1"
	_, err := r.db.Exec(ctx, deleteUserRequest, id)
	if err != nil {
		return err
	}

	return nil
}
