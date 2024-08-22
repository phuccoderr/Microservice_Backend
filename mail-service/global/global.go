package global

import (
	"github.com/segmentio/kafka-go"
	mail "github.com/xhit/go-simple-mail/v2"
	"mail-service/pkg/logger"
	"mail-service/pkg/setting"
)

var (
	Config  setting.Config
	Logger  *logger.LoggerZap
	Mailer  *mail.SMTPClient
	Consume *kafka.Reader
)
