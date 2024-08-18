package setting

type Config struct {
	Postgres PostgresSetting `mapstructure:"postgres"`
	Logger   LogSetting      `mapstructure:"logger"`
	Server   ServerSetting   `mapstructure:"server"`
	Kafka    KafkaSetting    `mapstructure:"kafka"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type PostgresSetting struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
	Sslmode  string `mapstructure:"sslmode"`
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
	Brokers string `mapstructure:"brokers"`
	Topic   string `mapstructure:"topic"`
}
