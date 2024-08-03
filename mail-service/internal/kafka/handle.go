package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"mail-service/internal/mail"
)

type VerifyCustomerMessage struct {
	Email     string `json:"email"`
	UrlVerify string `json:"url"`
}

func (c *Consumer) handleVerifyCustomer(message *sarama.ConsumerMessage) error {
	var verify VerifyCustomerMessage
	err := json.Unmarshal(message.Value, &verify)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return err
	}

	formMail := mail.Message{
		From:    "phuctapcode@gmail.com",
		To:      verify.Email,
		Subject: "Xác thực tài khoản",
		Body:    "Chẳng có gì cả",
		HTMLBody: fmt.Sprintf(`<h3>Vui lòng nhấn đường link dưới đây để xác thực tài khoản của bạn!</h3>
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
      </a>`, verify.UrlVerify),
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
