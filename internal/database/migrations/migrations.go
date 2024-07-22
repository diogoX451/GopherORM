package migrations

import (
	"fmt"

	"github.com/diogoX451/gopherORM/internal/database"
)

type Database = database.Database

type IMigration interface {
	Up() database.DatabaseTypes
	Down() database.DatabaseTypes
	GetTableName() string
}

type Migrations struct {
	db Database
}

func NewMigrations(db Database) *Migrations {
	return &Migrations{db: db}
}

func (m *Migrations) Run() error {
	fmt.Println("Running migrations")
	migrations := GetMigrations()
	fmt.Println("Migrations", migrations)
	for _, migration := range migrations {
		up := migration.Up()
		fmt.Println("Migration", migration.GetTableName(), "running")
		vf, err := m.db.TableExists(migration.GetTableName())
		if err != nil {
			return err
		}

		if !vf {
			if err := m.db.CreateTable(migration.GetTableName(), up); err != nil {
				return err
			}

			fmt.Println("Migration", migration.GetTableName(), "executed")
			m.db.Query("INSERT INTO migrations (table_name) VALUES ($1)", migration.GetTableName())
		}
	}

	return nil
}

// func (m *Migrations) Rollback() error {
// 	for _, migration := range migrations {
// 		if err := migration.Down(m.db); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

func (m *Migrations) Close() {
	m.db.Close()
}
