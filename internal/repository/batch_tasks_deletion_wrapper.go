package repository

import (
	"context"
	"sync"
	"time"
)

type BatchTaskDeletionSQLRepositoryWrapper struct {
	PostgreSQLTaskRepository
	batchChan chan struct{}
	once      sync.Once
}

func NewBatchTaskDeletionSQLRepositoryWrapper(wrapee PostgreSQLTaskRepository) *BatchTaskDeletionSQLRepositoryWrapper {
	return &BatchTaskDeletionSQLRepositoryWrapper{
		PostgreSQLTaskRepository: wrapee,
		batchChan:                make(chan struct{}, 10),
	}
}

func (w *BatchTaskDeletionSQLRepositoryWrapper) Delete(id ItemID) error {
	w.once.Do(func() {
		go w.observeDeletionBatch(w.batchChan)
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateTaskRequest := "UPDATE tasks SET deleted = true WHERE tid = $1"
	_, err := w.db.Exec(ctx, updateTaskRequest, id)
	if err != nil {
		return err
	}

	go func() {
		w.batchChan <- struct{}{}
	}()

	return nil
}

func (w *BatchTaskDeletionSQLRepositoryWrapper) observeDeletionBatch(batchChan <-chan struct{}) {
	for {
		batchLen := cap(batchChan)

		if len(batchChan) == batchLen {
			for i := 0; i < batchLen; i++ {
				<-batchChan
			}

			w.deleteAllMarkedTasks()
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func (w *BatchTaskDeletionSQLRepositoryWrapper) deleteAllMarkedTasks() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	deleteTasksRequest := "DELETE FROM tasks WHERE deleted = true"
	_, err := w.db.Exec(ctx, deleteTasksRequest)
	if err != nil {
		return err
	}

	return nil
}
