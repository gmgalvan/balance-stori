package business

import (
	"bytes"
	"context"
	"html/template"
	"math"
	"time"

	"github.com/stori/internal/schema"
)

const (
	//DateFormat  the layout format of the date strings in the CSV file
	DateFormat = "2006/01/02"
)

// Repository interface with repository logic functions
type Repository interface {
	InsertTransactionToDB(ctx context.Context, id, date, transaction string) error
	GetTransactionsFromBucket() ([]schema.Transaction, error)
}

// Service interface with services logic
type Service interface {
	SendEmail(subject, toAddress string, message []byte) error
}

// Business struct for stori
type Business struct {
	repo    Repository
	service Service
}

// NewBusiness creates a new Business instance
func NewBusiness(repo Repository, service Service) *Business {
	return &Business{
		repo:    repo,
		service: service,
	}
}

// CreateBalance calculates the balance and returns a summary
func (b *Business) CreateBalance() (schema.Summary, error) {
	var summary schema.Summary
	transactions, err := b.repo.GetTransactionsFromBucket()
	if err != nil {
		return summary, err
	}

	summary, err = b.calculateBalance(transactions)
	if err != nil {
		return summary, err
	}

	return summary, nil
}

// SendBalanceEmail sends email to specific address with balance
func (b *Business) SendBalanceEmail(subject, toAddress string, summary schema.Summary) error {
	message, err := b.createBalanceEmailMessage(summary)
	if err != nil {
		return err
	}
	err = b.service.SendEmail(subject, toAddress, message)
	if err != nil {
		return err
	}
	return nil
}

func (b *Business) calculateBalance(transactions []schema.Transaction) (schema.Summary, error) {
	var summary schema.Summary
	var balance float64
	totalCredit := 0.0
	totalDebit := 0.0
	numberCredit := 0
	numberDebit := 0
	summary.MonthTransactions = make(map[string]int64)
	for _, t := range transactions {
		balance = balance + t.Transaction
		date, err := time.Parse(DateFormat, t.Date)
		if err != nil {
			return summary, err
		}
		month := date.Month().String()
		summary.MonthTransactions[month] = summary.MonthTransactions[month] + 1
		if t.Transaction > 0 {
			totalCredit = totalCredit + t.Transaction
			numberCredit = numberCredit + 1
		}
		if t.Transaction < 0 {
			totalDebit = totalDebit + t.Transaction
			numberDebit = numberDebit + 1
		}
	}
	summary.TotalBalance = balance
	summary.AverageCreditAmount = math.Round((totalCredit/float64(numberCredit))*100) / 100
	summary.AverageDebitAmout = math.Round((totalDebit/float64(numberDebit))*100) / 100
	return summary, nil
}

func (b *Business) createBalanceEmailMessage(summary schema.Summary) ([]byte, error) {
	tmpl, err := template.ParseFiles("templates/balance.html")
	if err != nil {
		return []byte(""), err
	}
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, summary)
	if err != nil {
		return []byte(""), err
	}
	return []byte(tpl.String()), nil
}
