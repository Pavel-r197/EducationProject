package migrations

import (
	"context"
	"database/sql"
	"errors"
	"io/fs"
	"path"
)

type migration struct {
	Version  int64
	Name     string
	SQL      string
	CheckSum string
}

func UP(ctx context.Context, db *sql.DB) error {
	if db == nil {
		return errors.New("Ошибка миграции, подключение отсутствует")
	}

}

func DOWN(ctx context.Context, db *sql.DB) error {

}

func loadSQLMigration(source fs.FS) ([]migration, error) {
	entries, err := fs.ReadDir(source, "sql")
	if err != nil {
		return nil, err
	}
	result := make([]migration, 0, len(entries))
	versions := make(map[int64]string)
	for _, value := range entries {
		if value.IsDir() || path.Ext(value.Name()) != ".sql" {
			continue
		}

	}
}
