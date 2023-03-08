package transport

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stori/internal/schema"
)

type StoriUC interface {
	CreateBalance() (schema.Summary, error)
	SendBalanceEmail(subject, toAddress string, summary schema.Summary) error
}

// StoriTransport struct for transport
type StoriTransport struct {
	storiBusiness StoriUC
	cfgs          *schema.PopulatedConfigs
}

// StoriTransport creates new transport for stori
func NewStoriTransport(sc StoriUC, cfgs *schema.PopulatedConfigs) *StoriTransport {
	return &StoriTransport{
		storiBusiness: sc,
		cfgs:          cfgs,
	}
}

func (ts *StoriTransport) handleS3Event(ctx context.Context, event events.S3Event) error {
	summary, err := ts.storiBusiness.CreateBalance()
	if err != nil {
		log.Fatal(err)
	}

	subject := "Your Stori Financial statement"
	err = ts.storiBusiness.SendBalanceEmail(subject, ts.cfgs.EmailToSendBalance, summary)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (ts *StoriTransport) Start() {
	lambda.Start(ts.handleS3Event)
}
