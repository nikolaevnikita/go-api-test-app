package repository

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

func Migrate(dbDsn string, migratePath string) error {
	path := fmt.Sprintf("file://%s", migratePath)
	m, err := migrate.New(path, dbDsn)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return fmt.Errorf("failed to migrate db: %w", err)
	}

	return nil
}