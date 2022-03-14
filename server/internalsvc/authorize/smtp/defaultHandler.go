package smtp

import (
	"SoftwareDevelopment-Backend/config"
	"fmt"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"net/smtp"
	"net/textproto"
	"time"
)

type DefaultSMTP struct {
	log    *zap.Logger
	config *config.Config
	pool   *email.Pool
}

func (d *DefaultSMTP) SendCode(emailAddress string, code int) error {
	e := &email.Email{
		To:      []string{emailAddress},
		From:    fmt.Sprintf("%s <%s>", d.config.Services.Auth.SMTP.DisplayFrom, d.config.Services.Auth.SMTP.UserName),
		Subject: "PhotoLang Verification Code",
		Text:    []byte("Here's your verification code"),
		HTML:    []byte(fmt.Sprintf("<h1>Code: %d<h1>", code)),
		Headers: textproto.MIMEHeader{},
	}
	if err := d.pool.Send(e, 5*time.Second); err != nil {
		d.log.Error("sending email: ", zap.Error(err))
		return err
	}
	return nil
}

func InitDefaultSMTP(logger *zap.Logger, config *config.Config) *DefaultSMTP {
	p, err := email.NewPool(
		config.Services.Auth.SMTP.RemoteAddress,
		5,
		smtp.PlainAuth("", config.Services.Auth.SMTP.UserName, config.Services.Auth.SMTP.Password, config.Services.Auth.SMTP.Host),
	)
	if err != nil {
		logger.Fatal("initializing smtp service: ", zap.Error(err))
	}
	return &DefaultSMTP{
		log:    logger,
		config: config,
		pool:   p,
	}
}
