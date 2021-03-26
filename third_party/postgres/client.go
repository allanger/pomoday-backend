package database

import (
	"context"
	"fmt"

	// "github.com/allanger/pomoday-backend/internal/logger"
	"github.com/allanger/pomoday-backend/tools/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

var (
	log  = logger.NewLogger("database").Entry
	pool *pgxpool.Pool
	err  error
)

// Pool return db coonection pool
func Pool() *pgxpool.Pool {
	if pool == nil {
		log.Debug("Connection pool doesn't exist. Let's create a new one")
		err := OpenConnectionPool()
		if err != nil {
			// log.Error(err)
			return nil
		}
	}
	return pool
}

// OpenConnectionPool opens new connection pool
func OpenConnectionPool() error {
	log.Debug("Opening pool")
	pool, err = pgxpool.Connect(context.Background(), connectionString())
	if err != nil {
		return err
	}
	log.Debug(connectionString())
	return nil
}

func connectionString() string {
	dbUsername := viper.GetString("database_username")
	dbPassword := viper.GetString("database_password")
	dbName := viper.GetString("database_name")
	dbHost := viper.GetString("database_host")
	dbPort := viper.GetString("database_port")
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	return dbURI
}
