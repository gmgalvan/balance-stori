package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/mail"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

type emailData struct {
	Balance          string
	TransactionsJuly int
	TransactionsAug  int
	AvgDebit         string
	AvgCredit        string
}

func (s *Service) SendEmail(subject, toAddress string, message []byte) error {
	from := mail.Address{Address: s.confs.EmailAddressFrom}
	to := mail.Address{Address: toAddress}

	// Crear el mensaje MIME
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	var msg bytes.Buffer
	for k, v := range header {
		msg.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}

	msg.WriteString("\r\n" + base64.StdEncoding.EncodeToString(message))

	_, err := s.sesClient.SendRawEmail(&ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{Data: msg.Bytes()},
		Source:     aws.String(from.Address),
		Destinations: []*string{
			aws.String(to.Address),
		},
	})
	if err != nil {
		return err
	}
	return nil
}
