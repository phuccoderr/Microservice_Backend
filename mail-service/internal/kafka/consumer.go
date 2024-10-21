package kafka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"html/template"
	"log"
	"mail-service/global"
	"mail-service/internal/constants"
	"mail-service/internal/mail"
	"time"
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

	global.Logger.Info("Place order received", zap.Any("placeOrder", placeOrder))

	funcMap := template.FuncMap{
		"formatDate": formatDate,
		"calTotal":   calTotal,
	}

	tmpl, _ := template.New("place-order.html").Funcs(funcMap).ParseFiles("internal/kafka/templates/place-order.html")
	var body bytes.Buffer
	err = tmpl.Execute(&body, placeOrder)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	formMail := mail.Message{
		From:     constants.EMAIL_FROM,
		To:       placeOrder.CustomerId.Email,
		Subject:  constants.VERIFY_PLACEORDER,
		Body:     body.String(),
		HTMLBody: body.String(),
	}

	err = mail.SendSMTPMessage(formMail)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("Send mail %s failed", message.Topic), zap.Error(err))
		return err
	}
	return nil
}

func formatDate(t time.Time) string {
	return t.Format("02/01/2006 15:04:05") // dd/MM/yyyy HH:mm:ss format
}

func calTotal(price float64, sale float64, quantity int64) float64 {
	sale = price * (sale / 100)
	total := (price - sale) * float64(quantity)
	return total
}
