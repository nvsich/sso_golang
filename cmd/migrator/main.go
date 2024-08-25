package migrator

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
)

func main() {
	var (
		storagePath     string
		migrationsPath  string
		migrationsTable string
	)

	flag.StringVar(&storagePath, "storage-path", "", "Path to a directory containing storage backend files")
	flag.StringVar(&migrationsPath, "migrations-path", "", "Path to a directory containing migration files")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "Path to a table containing migration files")
	flag.Parse()

	if storagePath == "" {
		panic("storage-path is required")
	}

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)

	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations found")
			return
		}

		panic(err)
	}

	fmt.Println("migrated")
}