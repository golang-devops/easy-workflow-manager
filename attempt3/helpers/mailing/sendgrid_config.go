package mailing

type SendGridConfig struct {
	From struct {
		Name  string
		Email string
	}
	TemplateID string
}

func NewSendGridConfig(fromName, fromEmail, templateID string) (cfg *SendGridConfig) {
	cfg = &SendGridConfig{}
	cfg.From.Name = fromName
	cfg.From.Email = fromEmail
	cfg.TemplateID = templateID
	return
}
