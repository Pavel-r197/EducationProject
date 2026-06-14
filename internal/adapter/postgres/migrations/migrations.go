package migrations

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"path"
	"strconv"
	"strings"
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

// Читает sql файл из файловой системы
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
		// Мы получили 1, initial_scheme, nil
		version, name, err := parseFileName(value.Name())
		if err != nil {
			return nil, err
		}
		// {1:initial_scheme}
		if _, ok := versions[version]; ok {
			return nil, fmt.Errorf("Dublicate migration version %d", version)
		}
		content, err := fs.ReadFile(source, path.Join("sql", value.Name()))
		if err != nil {
			return nil, fmt.Errorf("Read migration %q:%w", value.Name(), err)
		}
		str := string(content)
		if str == "" {
			return nil, fmt.Errorf("Migration %q is empty", value.Name())
		}
		//дописать рассчитать хэш, добавить в мапу версию
	}
}

// Преобразует файл 000001_initial_scheme.sql в версию 1 и имя initial_scheme
func parseFileName(filename string) (int64, string, error) {
	base := strings.TrimSuffix(filename, path.Ext(filename))
	// 1-0001 , 2-initial_scheme, 3-true
	versionText, name, ok := strings.Cut(base, "_")
	if !ok || versionText == "" || name == "" {
		return 0, "", fmt.Errorf("Invalid migration filename %q", filename)
	}
	version, err := strconv.ParseInt(versionText, 10, 64)
	if err != nil || version <= 0 {
		return 0, "", fmt.Errorf("Invalid migration filename %q", filename)
	}
	return version, name, nil
}
