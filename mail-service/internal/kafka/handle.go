package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"mail-service/config"
	"mail-service/constants"
	"mail-service/internal/mail"
)

type ConsumerGroupHandler struct {
	Config *config.Config
}

func (c *ConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

type VerifyCustomerMessage struct {
	Email     string `json:"email"`
	UrlVerify string `json:"url"`
}

func (c *ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		switch msg.Topic {
		case "customer-verify-events-topics":
			c.handleVerifyCustomer(msg)
		case "customer-password-events-topics":
			c.handleVerifyPassword(msg)
		default:
			fmt.Println("Unknown topic")
		}
		sess.MarkMessage(msg, "")
	}
	return nil
}

func (c *ConsumerGroupHandler) handleVerifyPassword(message *sarama.ConsumerMessage) error {
	var verify VerifyCustomerMessage
	err := json.Unmarshal(message.Value, &verify)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return err
	}
	formMail := mail.Message{
		From:    constants.EMAIL_FROM,
		To:      verify.Email,
		Subject: constants.CHANGE_PASSWORD,
		Body:    constants.BODY,
		HTMLBody: fmt.Sprintf(`<h3>%s</h3>
      				<a href="%s">
			<button
			  style="
				padding: 20px;
				background-color: gray;
				border-radius: 10px;
				color: white;
			  "
			>
			  Change Password
			</button>
      </a>`, verify.UrlVerify, constants.TITLE_RESET_PASSWORD),
	}

	mailer := &mail.Mailer{
		Config: c.Config,
	}
	err = mailer.SendSMTPMessage(formMail)
	if err != nil {
		log.Printf("Error send email: %v", err)
		return err
	}
	return nil
}

func (c *ConsumerGroupHandler) handleVerifyCustomer(message *sarama.ConsumerMessage) error {
	var verify VerifyCustomerMessage
	err := json.Unmarshal(message.Value, &verify)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return err
	}

	formMail := mail.Message{
		From:    constants.EMAIL_FROM,
		To:      verify.Email,
		Subject: constants.VERIFY_ACCOUNT,
		Body:    constants.BODY,
		HTMLBody: fmt.Sprintf(`<h3>%s</h3>
      				<a href="%s">
			<button
			  style="
				padding: 20px;
				background-color: gray;
				border-radius: 10px;
				color: white;
			  "
			>
			  Verify customer
			</button>
      </a>`, verify.UrlVerify, constants.TITLE_VERIFY_ACCOUNT),
	}

	mailer := &mail.Mailer{
		Config: c.Config,
	}
	err = mailer.SendSMTPMessage(formMail)
	if err != nil {
		log.Printf("Error send email: %v", err)
		return err
	}
	return nil
}
