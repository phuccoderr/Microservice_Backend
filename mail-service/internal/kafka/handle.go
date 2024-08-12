package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"mail-service/config"
	"mail-service/constants"
	"mail-service/internal/mail"
	"time"
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

type PlaceOrderMessage struct {
	ProductTotal  float64   `json:"product_total"`
	ProductCost   float64   `json:"product_cost"`
	ShippingCost  float64   `json:"shipping_cost"`
	Address       string    `json:"address"`
	DeliverDays   time.Time `json:"deliver_days"`
	PaymentMethod string    `json:"payment_method"`
	CustomerEmail string    `json:"customer_email"`
}

func (c *ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		switch msg.Topic {
		case constants.KAFKA_TOPIC_VERIFY_ACCOUNT:
			c.handleVerifyCustomer(msg)
		case constants.KAFKA_TOPIC_VERIFY_PASSWORD:
			c.handleVerifyPassword(msg)
		case constants.KAFKA_TOPIC_PLACE_ORDER:
			c.handlePlaceOrderCustomer(msg)
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
      </a>`, constants.TITLE_RESET_PASSWORD, verify.UrlVerify),
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
      </a>`, constants.TITLE_VERIFY_ACCOUNT, verify.UrlVerify),
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

func (c *ConsumerGroupHandler) handlePlaceOrderCustomer(message *sarama.ConsumerMessage) error {
	var placeOrder PlaceOrderMessage
	err := json.Unmarshal(message.Value, &placeOrder)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return err
	}

	formMail := mail.Message{
		From:     constants.EMAIL_FROM,
		To:       placeOrder.CustomerEmail,
		Subject:  constants.VERIFY_PLACEORDER,
		Body:     constants.BODY,
		HTMLBody: fmt.Sprintf(`<h3>%s</h3>`, constants.TITLE_PLACE_ORDER),
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
