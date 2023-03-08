package services

import (
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/stori/internal/schema"
)

type Service struct {
	sesClient *ses.SES
	confs     *schema.PopulatedConfigs
}

func NewService(sesClient *ses.SES, cfgs *schema.PopulatedConfigs) *Service {
	return &Service{
		sesClient: sesClient,
		confs:     cfgs,
	}
}
