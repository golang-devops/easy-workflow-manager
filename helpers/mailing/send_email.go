package mailing

import (
	"errors"
	"fmt"
	nativeMail "net/mail"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/golang-devops/easy-workflow-manager/logging"
)

func SendHTMLEmail(logger logging.Logger, cfg *SendGridConfig, subject, htmlBody string, recipient *nativeMail.Address, substitutions map[string]string) error {
	from := mail.NewEmail(cfg.From.Name, cfg.From.Email)
	recipientAddress := mail.NewEmail(recipient.Name, recipient.Address)
	content := mail.NewContent("text/html", htmlBody)

	email := mail.NewV3MailInit(from, subject, recipientAddress, content)

	//TODO: Might want to be able to specify this or pass it in
	templateID := cfg.TemplateID
	email.SetTemplateID(templateID)

	for key, val := range substitutions {
		email.Personalizations[0].SetSubstitution(key, val)
	}

	request := SendgridRequestFactory.Request()
	request.Method = "POST"
	request.Body = mail.GetRequestBody(email)
	resp, err := sendgrid.API(*request)
	if err != nil {
		userMessage := "Failed to send Email"
		logger.WithError(err).WithField("email-subject", subject).Error(userMessage)
		return errors.New(userMessage)
	}

	if resp.StatusCode != 202 {
		userMessage := "Email sending failed with invalid Status"
		logger.WithField("email-subject", subject).Error(fmt.Sprintf("Sendgrid request failed. StatusCode = %d. Response body: %s", resp.StatusCode, resp.Body))
		return errors.New(userMessage)
	}

	return nil
}
