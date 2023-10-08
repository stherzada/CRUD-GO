package connections

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"CRUD-GO/db"
)

type DatabaseConnection struct {
	DB *sql.DB
	Q  *db.Queries
}

func (d *DatabaseConnection) Close() error {
	return d.DB.Close()
}

func NewPostgresConnection() (*DatabaseConnection, error) {
	conn, err := sql.Open(
		"postgres",
		"postgres://crud-go:crud-go@localhost:5432/crud-go?sslmode=disable",
	)
	if err != nil {
		return nil, err
	}

	db := db.New(conn)
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance("file://sql/migrations", "postgres", driver)
	if err != nil {
		return nil, err
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	return &DatabaseConnection{conn, db}, nil
}
