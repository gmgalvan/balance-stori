package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stori/internal/schema"
)

// Repostory struct for entities repository layer
type Repostory struct {
	dbclient *sql.DB
	s3client *s3.S3
	configs  *schema.PopulatedConfigs
}

// NewRepostory returns the repo struct
func NewRepostory(client *sql.DB, s3client *s3.S3, configs *schema.PopulatedConfigs) *Repostory {
	return &Repostory{
		dbclient: client,
		s3client: s3client,
		configs:  configs,
	}
}

// rollBackTransactionOnError ensures roll back transaction
func rollBackTransactionOnError(ctx context.Context, tx *sql.Tx, rootError error) {
	errRollBack := tx.Rollback()
	if errRollBack != nil {
		log.Printf("unable to rollback transaction on %v, due to: %v", rootError, errRollBack)
	}
}
