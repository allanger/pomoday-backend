package migrations

const createChainsTable = `
CREATE TABLE IF NOT EXISTS tasks (
    userid UUID REFERENCES users(id) ON DELETE CASCADE,
		id INTEGER NOT NULL,
		uuid UUID  PRIMARY KEY,
		archived BOOLEAN,
		tag TEXT NOT NULL,
		title TEXT NOT NULL,
		status INTEGER NOT NULL,
		lastaction BIGINT NOT NULL,
		logs JSONB
);
`

func init() {
	migration := &migration{
		name:    "Create tasks table",
		version: "0.0.3",
		up:      createChainsTable,
		down:    "DROP TABLE IF EXISTS tasks",
	}
	newMigration(migration)
}
