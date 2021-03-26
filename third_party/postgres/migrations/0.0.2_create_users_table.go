package migrations

const createUsersTable = `
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT
);
`

func init() {
	migration := &migration{
		name:    "Create users table",
		version: "0.0.2",
		up:      createUsersTable,
		down:    "DROP TABLE IF EXISTS users",
	}
	newMigration(migration)
}
