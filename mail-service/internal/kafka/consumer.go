package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"mail-service/global"
	"mail-service/internal/constants"
	"mail-service/internal/mail"
)

func HandleVerifyCustomer(message *kafka.Message) error {
	var verify VerifyCustomerMessage
	err := json.Unmarshal(message.Value, &verify)
	if err != nil {
		global.Logger.Error("Unmarshal json failed", zap.Error(err))
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

	err = mail.SendSMTPMessage(formMail)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("Send mail %s failed", message.Topic), zap.Error(err))
		return err
	}
	return nil
}

func HandleVerifyPassword(message *kafka.Message) error {
	var verify VerifyCustomerMessage
	err := json.Unmarshal(message.Value, &verify)
	if err != nil {
		global.Logger.Error("Unmarshal json failed", zap.Error(err))
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

	err = mail.SendSMTPMessage(formMail)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("Send mail %s failed", message.Topic), zap.Error(err))
		return err
	}
	return nil
}

func HandlePlaceOrderCustomer(message *kafka.Message) error {
	var placeOrder PlaceOrderMessage
	err := json.Unmarshal(message.Value, &placeOrder)
	if err != nil {
		global.Logger.Error("Unmarshal json failed", zap.Error(err))
		return err
	}

	formMail := mail.Message{
		From:     constants.EMAIL_FROM,
		To:       placeOrder.CustomerEmail,
		Subject:  constants.VERIFY_PLACEORDER,
		Body:     constants.BODY,
		HTMLBody: fmt.Sprintf(`<h3>%s</h3>`, constants.TITLE_PLACE_ORDER),
	}

	err = mail.SendSMTPMessage(formMail)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("Send mail %s failed", message.Topic), zap.Error(err))
		return err
	}
	return nil
}
