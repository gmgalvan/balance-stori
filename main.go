package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/stori/internal/business"
	"github.com/stori/internal/database"
	"github.com/stori/internal/repository"
	"github.com/stori/internal/services"
	"github.com/stori/internal/transport"

	"github.com/stori/internal/configs"
)

func main() {
	ctx := context.Background()

	cfg, err := configs.PopulateConfigs()
	if err != nil {
		log.Fatalf("Could not load configurations. Error: [%s]", err)
	}

	// AWS session configs
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(cfg.AWSAccessToken, cfg.AWSSecret, ""),
	})
	if err != nil {
		log.Fatal(err)
	}

	// start database configuration
	db, err := database.NewClientDB(ctx, "postgres", cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Run migrations using goose framework https://github.com/pressly/goose
	m := database.NewMigration(db, cfg.MigrationsPath)
	err = m.StartMigration()
	if err != nil {
		log.Fatalf("Migration failed. Error: [%s]", err)
	}

	// configure email service
	service := services.NewService(ses.New(sess), cfg)

	// Repo layer
	repo := repository.NewRepostory(db, s3.New(sess), cfg)

	// business cases layer
	businessCases := business.NewBusiness(repo, service)

	// transport layer
	tr := transport.NewStoriTransport(businessCases, cfg)
	tr.Start()
}
