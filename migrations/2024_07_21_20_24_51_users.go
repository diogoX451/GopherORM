package migrations

import (
	"github.com/diogoX451/gopherORM/internal/database"
	"github.com/diogoX451/gopherORM/internal/database/migrations"
)

var _ migrations.IMigration = (*Users)(nil)

func (m Users) GetTableName() string {
	return "users"
}

func init() {
	migrations.RegisterMigration(NewUsers())
}

type Users struct {
}

func NewUsers() *Users {
	return &Users{}
}

func (m Users) Up() database.DatabaseTypes {
	var db database.DatabaseTypes
	return *database.NewDatabaseTypes(
		db.Id("id"),
		db.String("name", 255),
		db.String("email", 255),
		db.String("password", 255),
		db.Timestamp(),
	)
}

func (m Users) Down() database.DatabaseTypes {
	return *database.NewDatabaseTypes()
}
