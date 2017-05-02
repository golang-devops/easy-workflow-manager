package mailing

import (
	"github.com/sendgrid/rest"
	sendgrid "github.com/sendgrid/sendgrid-go"
)

var SendgridRequestFactory *sendgridRequestFactory

type sendgridRequestFactory struct {
	apiKey   string
	endpoint string
	host     string
}

func (s *sendgridRequestFactory) Request() *rest.Request {
	req := sendgrid.GetRequest(s.apiKey, s.endpoint, s.host)
	return &req
}

func InitSendgridRequestFactory(apiKey string) {
	SendgridRequestFactory = &sendgridRequestFactory{
		apiKey:   apiKey,
		endpoint: "/v3/mail/send",
		host:     "https://api.sendgrid.com",
	}
}
