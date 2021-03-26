package migrations

import (
	"context"

	database "github.com/allanger/pomoday-backend/third_party/postgres"
	"github.com/allanger/pomoday-backend/tools/logger"
	"github.com/spf13/viper"
)

var (
	migrations     []*migration
	log            = logger.NewLogger("migrations").Entry
	currentVersion string
	currentName    string
	wishedVersion  string
)

var (
	getCurrentVersion = "SELECT version FROM db_version;"
	setVersion        = "INSERT INTO db_version (version, last_migration) VALUES ($1, $2);"
)

type migration struct {
	name    string
	version string
	up      string
	down    string
}

func newMigration(migration *migration) {
	migrations = append(migrations, migration)
}

func getLatestVersion() string {
	return migrations[len(migrations)-1].version

}

// Migrate runs go-pg migrations
func Migrate() {
	ctx := context.Background()

	if err := database.OpenConnectionPool(); err != nil {
		log.Fatal(err)
	}

	log.Info("Starting migrations")
	pool := database.Pool()

	err := pool.QueryRow(ctx, getCurrentVersion).Scan(&currentVersion)
	if err != nil {
		currentVersion = "0.0.0"
	}

	if viper.Get("database_version") != nil {
		log.Info("Migrating to version ", viper.Get("database_version"))
		wishedVersion = viper.GetString("database_version")
	} else {
		log.Info("Migrating to the latest version")
		wishedVersion = getLatestVersion()
	}

	if currentVersion >= wishedVersion {
		log.Info("Already on wished database version")
		return
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback(ctx)

	for _, migration := range migrations {
		if migration.version > wishedVersion {
			break
		}
		_, err = tx.Exec(ctx, migration.up)
		if err != nil {
			log.Fatal(err)
		}
		currentVersion = migration.version
		currentName = migration.name
	}
	_, err = tx.Exec(ctx, setVersion, currentVersion, currentName)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
