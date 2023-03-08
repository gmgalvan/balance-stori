package database

import (
	"database/sql"
	// Driver for postgres
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

// MigrationStore struct for a goose migration
type MigrationStore struct {
	db           *sql.DB
	dirMigration string
}

// NewMigration start a new goose implementation
func NewMigration(client *sql.DB, dirMigration string) *MigrationStore {
	return &MigrationStore{
		db:           client,
		dirMigration: dirMigration,
	}
}

// StartMigration starts a new migration version going forward
func (m MigrationStore) StartMigration() error {
	return m.MigrationUp()
}

// MigrationUp function for version forward migration
func (m MigrationStore) MigrationUp() error {
	err := goose.Up(m.db, m.dirMigration)
	if err != nil {
		return err
	}
	return goose.Status(m.db, m.dirMigration)
}
