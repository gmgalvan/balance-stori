package database

import (
	"context"
	"fmt"
	"net/url"

	// Driver for postgres
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/stori/internal/schema"
)

// NewClientDB create a new connection to the postgres server
func NewClientDB(ctx context.Context, dbDriver string, cfg *schema.PopulatedConfigs) (*sql.DB, error) {
	// start db connection with postgres
	pgURL := url.URL{
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%d", cfg.DataBaseHost, cfg.DatabasePort),
		User:   url.UserPassword(cfg.DatabaseUser, cfg.DatabasePassworld),
		Path:   cfg.DatabaseName,
	}

	options := url.Values{}
	options.Set("sslmode", cfg.SSLDatabaseMode)
	pgURL.RawQuery = options.Encode()

	connDB := pgURL.String()

	db, err := sql.Open(dbDriver, connDB)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
