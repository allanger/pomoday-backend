package migrations

const createDBVersionTable = `
	CREATE TABLE IF NOT EXISTS db_version (
		version TEXT PRIMARY KEY,
		last_migration TEXT
);
`

func init() {
	migration := &migration{
		name:    "Create db_version table",
		version: "0.0.1",
		up:      createDBVersionTable,
		down:    "DROP TABLE IF EXISTS db_version",
	}
	newMigration(migration)
}
