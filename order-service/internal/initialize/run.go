package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitPostgres()
	IniteKafka()
	InitRedis()
	r := InitRouter()

	r.Run(":9170")
}
