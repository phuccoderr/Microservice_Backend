package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Logger LogSetting    `mapstructure:"logger"`
	Jwt    JwtSetting    `mapstructure:"jwt"`
	Redis  RedisSetting  `mapstructure:"redis"`
	Kafka  KafkaSetting  `mapstructure:"kafka"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type JwtSetting struct {
	Secret string `mapstructure:"secret"`
	Expire int    `mapstructure:"expire"`
}

type LogSetting struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}

type RedisSetting struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type KafkaSetting struct {
	Brokers string `mapstructure:"brokers"`
	Topic   string `mapstructure:"topic"`
}
