package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v5/pgxpool"

	"simple_go/config"
	"simple_go/database/db"
	"simple_go/http"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
)

func main() {

	// configaration
	log.Info().Msg("Starting the application")
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal().Err(err).Msg("configuration loading failed")
	}
	gin.SetMode(config.GinMode)

	//structured logging
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	if config.GinMode == "debug" {
		log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()

	}
	// initialize the database connection pool
	log.Info().Msg("Creating a new connection pool")
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("connection pool creation failed")
	}
	defer connPool.Close()

	// test the database connection
	log.Info().Msg("Pinging the database")
	err = connPool.Ping(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("database ping failed")
	}

	// database migration
	log.Info().Msg("Applying database migrations")
	m, err := migrate.New(config.MigrationURL, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("migration failed")
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("database migration failed")
	}

	store := db.NewStore(connPool)
	server := http.NewServer(store, &config)

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
