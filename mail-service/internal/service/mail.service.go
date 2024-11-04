package service

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"html/template"
	"log"
	"mail-service/global"
	"mail-service/internal/constants"
	"mail-service/internal/dto"
	"mail-service/internal/mail"
)

type IMailService interface {
	Feedback(feedback dto.Feedback) error
}

type MailService struct {
}

func (m MailService) Feedback(feedback dto.Feedback) error {
	tmpl, _ := template.New("feedback.html").ParseFiles("internal/kafka/templates/feedback.html")
	var body bytes.Buffer
	err := tmpl.Execute(&body, feedback)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	formMail := mail.Message{
		From:     feedback.Email,
		To:       constants.EMAIL_FROM,
		Subject:  constants.FEEDBACK_CUSTOMER,
		Body:     body.String(),
		HTMLBody: body.String(),
	}

	err = mail.SendSMTPMessage(formMail)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("Send mail %s failed", feedback.Email), zap.Error(err))
		return err
	}
	return nil
}

func NewMailService() IMailService {
	return &MailService{}
}
