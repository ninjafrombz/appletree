//Filename: cmd/api/main.go

package main

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"time"

	"appletree.desireamagwula.net/internals/data"
	"appletree.desireamagwula.net/internals/jsonlog"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

// The configuration settings

type config struct {
	port int
	env  string // development, staging, production, etc.
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	limiter struct {
		rps     float64 // requests/second
		burst   int
		enabled bool
	}
}

// DEpendency injection
type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
}

func main() {
	var cfg config
	// read in the flags that are needed to populate our config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development | staging | production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("APPLETREE_DB_DSN"), "Postgresql DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "Postgresql max open CONNECTIONS")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "Postgresql idle open CONNECTIONS")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgresQL max connection idle time")
	// These are flags for the rate limiter
	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.Parse()

	// create a logger
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	// CReate the connection pool
	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	defer db.Close()
	// LOg the succesful
	logger.PrintInfo("database connection pool established", nil)

	//create an instancr of application struct
	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	// Call app.serve to start the server
	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}

// The openDB funtion returns a *sql.DB connection pool

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)
	// Create a context with a 5-second timeout timeline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
