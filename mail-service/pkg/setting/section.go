package setting

type Config struct {
	Server  ServerSetting  `mapstructure:"server"`
	Logger  LogSetting     `mapstructure:"logger"`
	Kafka   KafkaSetting   `mapstructure:"kafka"`
	Mailhog MailhogSetting `mapstructure:"mailhog"`
}

type ServerSetting struct {
	Port int `mapstructure:"port"`
}

type LogSetting struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}

type KafkaSetting struct {
	Brokers string   `mapstructure:"brokers"`
	Topics  []string `mapstructure:"topics"`
	Group   string   `mapstructure:"group"`
}

type MailhogSetting struct {
	MailDomain     string `mapstructure:"mail_domain"`
	MailHost       string `mapstructure:"mail_host"`
	MailPort       int    `mapstructure:"mail_port"`
	MailUsername   string `mapstructure:"mail_username"`
	MailPassword   string `mapstructure:"mail_password"`
	MailEncryption string `mapstructure:"mail_encryption"`
}
