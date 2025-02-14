package migratedb

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(root_path string) {

	m, err := migrate.New(
		"file://"+root_path+"/db/migrations",
		"postgres://postgres:password@localhost:5432/fiber?sslmode=disable")
	if err != nil {
		fmt.Println("error:", err)
	}
	if err := m.Up(); err != nil {
		fmt.Println("error up:", err)
	}
}
