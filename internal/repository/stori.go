package repository

import (
	"context"
	"encoding/csv"
	"io"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stori/internal/schema"
)

func (repo *Repostory) InsertTransactionToDB(ctx context.Context, id, date, transaction string) error {
	createRecord := "INSERT INTO transactions (id, date, transaction) VALUES ($1, $2, $3)"
	tx, err := repo.dbclient.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(createRecord)
	if err != nil {
		rollBackTransactionOnError(ctx, tx, err)
		return err
	}
	defer stmt.Close()

	return err
}

func (repo *Repostory) GetTransactionsFromBucket() ([]schema.Transaction, error) {
	var transactions []schema.Transaction

	resp, err := repo.s3client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(repo.configs.S3Bucket),
		Key:    aws.String(repo.configs.TransactionsFile),
	})
	if err != nil {
		return transactions, err
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)

	reader.Comma = ','
	reader.Comment = '#'
	reader.FieldsPerRecord = -1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return transactions, err
		}

		id, _ := strconv.Atoi(record[0])
		transaction, _ := strconv.ParseFloat(record[2], 64)

		t := schema.Transaction{
			ID:          id,
			Date:        record[1],
			Transaction: transaction,
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
