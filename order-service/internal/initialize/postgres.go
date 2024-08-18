package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"order-service/global"
	"order-service/internal/models"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(errString)
	}
}

func InitPostgres() {
	m := global.Config.Postgres

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		m.Host, m.Username, m.Password, m.Dbname, m.Port, m.Sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	checkErrorPanic(err, "InitPostgres error")

	global.Logger.Info("InitPostgres success")
	global.Mdb = db

	migrateTable()
}

func migrateTable() {
	err := global.Mdb.AutoMigrate(
		&models.Order{},
		&models.OrderDetails{},
	)
	checkErrorPanic(err, "migrate table error")
}
