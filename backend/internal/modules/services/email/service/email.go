package service

import (
	"context"
	"crypto/tls"
	_ "embed"
	"fmt"
	"gopkg.in/gomail.v2"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/services/email/dto"
	"waterfall-backend/internal/utils/query"
)

const (
	gmailSmtpHost = "smtp.gmail.com"
	gmailSmtpPort = 587
)

type emailLetter struct {
	toEmail  string
	subject  string
	htmlBody string
}

type EmailService struct {
	SystemEmailAddress  string
	SystemEmailPassword string
	DevMode             bool
	FrontendUrl         string
}

func NewEmailService(cfg config.Config) *EmailService {
	return &EmailService{
		SystemEmailAddress:  cfg.SystemEmailAddress,
		SystemEmailPassword: cfg.SystemEmailPassword,
		DevMode:             cfg.DevMode,
		FrontendUrl:         cfg.FrontendUrl,
	}
}

// SendVerificationEmail отправить письмо со ссылкой для подтверждения почты пользователя
func (e *EmailService) SendVerificationEmail(ctx context.Context, toEmail string, dtm *dto.Verification) error {
	html, err := parseTemplate(Verification, e.makeVerificationModel(dtm))
	if err != nil {
		return err
	}

	letter := emailLetter{
		toEmail:  toEmail,
		subject:  "Добро пожаловать!",
		htmlBody: html,
	}
	return e.sendEmail(ctx, letter)
}

func (e *EmailService) makeVerificationModel(dtm *dto.Verification) *VerificationModel {
	opts := VerificationOptions{Token: dtm.VerificationToken}
	queryParams := query.Encode(opts)

	return &VerificationModel{
		Link: fmt.Sprintf("%s/redirect/verify-email?%s", e.FrontendUrl, queryParams),
	}
}

// sendEmail отправляет письмо
func (e *EmailService) sendEmail(_ context.Context, letter emailLetter) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", e.SystemEmailAddress)
	msg.SetHeader("To", letter.toEmail)
	msg.SetHeader("Subject", letter.subject)
	msg.SetBody("text/html", letter.htmlBody)

	dealer := gomail.NewDialer(
		gmailSmtpHost,
		gmailSmtpPort,
		e.SystemEmailAddress,
		e.SystemEmailPassword,
	)

	if e.DevMode {
		dealer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	err := dealer.DialAndSend(msg)
	if err != nil {
		return err
	}

	return nil
}
