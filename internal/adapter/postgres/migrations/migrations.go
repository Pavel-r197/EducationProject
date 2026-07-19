package migrations

import (
	"context"
	"crypto"
	"crypto/sha256"
	"database/sql"
	"embed"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"path"
	"strconv"
	"strings"
)

const createSchemeMigrations = `
CREATE TABLE IF NOT EXISTS scheme_migrations (
    version BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    checksum TEXT NOT NULL,
    applied_at timestamptz NOT NULL DEFAULT NOW()
)
`

const deleteSchemeMigrations = ``

//go:embed sql/*.sql
var sqlfiles embed.FS

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
	if ctx == nil {
		ctx = context.Background()
	}
	if _, err := db.ExecContext(ctx, createSchemeMigrations); err != nil {
		return fmt.Errorf("Create scheme migrations table: %w", err)
	}
	migrations, err := loadSQLMigration(sqlfiles) {
		if err != nil{
			return err
		}
	}
	applied, err := loadAppliedMigrations(ctx, db)
	if err != nil {
		return err
	}
	for _, item := range migrations {
		if checksum, ok := applied[item.Version]; ok {
			if item.CheckSum != checksum {
				return fmt.Errorf("Migration %d checksum mismatch", item.Version)
			}
		}
		continue
	}
    // осталось выполнить миграции, остановились тут, будем открывать транзакцию и выплнять sql запрос и вставлять в схему миграции наше новое значение и комитить.
}

func DOWN(ctx context.Context, db *sql.DB) error {
	if db == nil {
		return errors.New("Ошибка миграции, подключение отсутствует")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if _, err := db.ExecContext(ctx, deleteSchemeMigrations); err != nil {
		return fmt.Errorf("Delete scheme migrations table: %w", err)
	}
	migrations, err := loadSQLMigration(sqlfiles) {
		if err != nil{
			return err
		}
	}

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
		checkSum := calculateCheckSun(str)
		versions[version] = value.Name()
		m := migration{Version: version, Name: name, SQL: str, CheckSum: checkSum}
		result = append(result, m)
	}
	return result, err
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

func calculateCheckSun(sqltext string) string {
	sum := sha256.Sum256([]byte(sqltext))
	return hex.EncodeToString(sum[:])
}

func loadAppliedMigrations (ctx context.Context, db *sql.DB) (map[int64]string, error) {
	q := "SELECT version, checksum FROM scheme_migrations"
	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("Load applied migrations: %w", err)
	}
	defer rows.Close()
	applied := make(map[int64]string)
	for rows.Next() {
		var version int64
		var checksum string
		if err := rows.Scan(&version, &checksum); err != nil {
			return nil, fmt.Errorf("Scan applied migrations: %w", err)
		}
		applied[version] = checksum
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Iterate applied migrations: %w", err)
	}
	return applied, nil
}