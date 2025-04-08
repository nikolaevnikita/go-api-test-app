package repository

import (
	"github.com/nikolaevnikita/go-api-test-app/internal/domain/models"

	"context"
	"time"
	"github.com/jackc/pgx/v5"
)

type PostgreSQLTaskRepository struct {
	db *pgx.Conn
}

// MARK: Fabric

func NewPostgreSQLTaskRepository(ctx context.Context, connString string) (*PostgreSQLTaskRepository, error) {
	db, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	taskRepository := PostgreSQLTaskRepository {
		db: db,
	}

	return &taskRepository, nil
}

// MARK: CRUD operations

func (r *PostgreSQLTaskRepository) Get(id ItemID) (*models.Task, error) {
	// при отстутсвии айдишника в бд - должна выдавать ошибку
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	row := r.db.QueryRow(ctx, "SELECT * from tasks WHERE tid = $1", id)

	var task models.Task
	err := row.Scan(&task.TID, &task.Title, &task.Description, &task.Status)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *PostgreSQLTaskRepository) GetAll() ([]*models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks []*models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.TID, &task.Title, &task.Description, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (r *PostgreSQLTaskRepository) Create(item models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	insertTaskRequest := "INSERT INTO tasks (tid, title, description, status) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(ctx, insertTaskRequest, item.TID, item.Title, item.Description, item.Status)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgreSQLTaskRepository) Update(id ItemID, item models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	updateTaskRequest := "UPDATE tasks SET title = $2, description = $3, status = $4 WHERE tid = $1"
	_, err := r.db.Exec(ctx, updateTaskRequest, id, item.Title, item.Description, item.Status)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgreSQLTaskRepository) Delete(id ItemID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	deleteTaskRequest := "DELETE FROM tasks WHERE tid = $1"
	_, err := r.db.Exec(ctx, deleteTaskRequest, id)
	if err != nil {
		return err
	}

	return nil
}
