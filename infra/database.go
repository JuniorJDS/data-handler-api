package infra

import (
	"context"
	"log"

	configs "github.com/JuniorJDS/data-handler-api/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func ConnectDB() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(configs.GetSettings()["POSTGRES"])
	if err != nil {
		log.Panicf("Failed to create config to database connection: %v\n", err)
	}

	config.MinConns = 1
	config.MaxConns = 150

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Panicf("Failed to connectin database: %v\n", err)
	}
	return conn

}

func GetDB() *pgxpool.Pool {
	if db == nil {
		db = ConnectDB()
	}

	return db
}
