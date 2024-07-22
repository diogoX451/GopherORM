package providers

import (
	"github.com/diogoX451/gopherORM/internal/database"
)

type DatabaseProviders struct {
	database database.Database
}

func NewDatabaseProviders() *DatabaseProviders {
	return &DatabaseProviders{}
}

func (d *DatabaseProviders) Connect() database.Database {
	db := database.InitializeDatabaseFactory()
	db.SetCloseAutomaticConn(0)
	db.SetConnection(10)
	db.SetMinConnections(5)

	connect := db.Connect()

	if connect != nil {
		return db
	}

	// mg := migrations.NewMigrations(db)
	// mg.Run([]migrations.IMigration{
	// 	querys.NewInit(),
	// 	querys.NewHistorys(),
	// })

	return db

}
