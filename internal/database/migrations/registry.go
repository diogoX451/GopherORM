package migrations

type MigrationRegistry struct {
	migrations []IMigration
}

var registry = &MigrationRegistry{}

func RegisterMigration(migration IMigration) {
	registry.migrations = append(registry.migrations, migration)
}

func GetMigrations() []IMigration {
	return registry.migrations
}
