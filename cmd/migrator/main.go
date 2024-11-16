package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var dbHost, dbPort, dbUser, dbPassword, dbName, migrationsPath, migrationsTable string

	flag.StringVar(&dbHost, "db-host", "db", "database host")

	flag.StringVar(&dbPort, "db-port", "5432", "database port")

	flag.StringVar(&dbUser, "db-user", "test", "database user")

	flag.StringVar(&dbPassword, "db-password", "1234", "database password")

	flag.StringVar(&dbName, "db-name", "db1", "database name")

	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")

	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse()

	if dbUser == "" || dbPassword == "" || dbName == "" {
		panic("database user, password and name are required")
	}
	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	postgresURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&x-migrations-table=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, migrationsTable)

	m, err := migrate.New(
		"file://"+migrationsPath,
		postgresURL,
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}

		panic(err)
	}
	fmt.Println("migrations applied")
}

type Log struct {
	verbose bool
}

func (l *Log) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (l *Log) Verbose() bool {
	return false
}
