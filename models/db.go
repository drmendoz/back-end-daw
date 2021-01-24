package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/drmendoz/backend-gin-gorm/utils"
)

var Db *gorm.DB

func init() {
	user := utils.Viper.Get("DB_USER").(string)
	password := utils.Viper.Get("DB_PASS").(string)
	server := utils.Viper.Get("DB_SERVER").(string)
	database := utils.Viper.Get("DB_NAME").(string)
	port := utils.Viper.GetString("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, server, port, database)
	utils.Log.Info(dsn)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Log.Fatal("Error al conectar base de datos", err)
	}
	migrarTablas()
}

func migrarTablas() {
	err := Db.AutoMigrate(&Administrador{})
	if err != nil {
		utils.Log.Fatal("Error al migrar modelos")
	}
}

type Tabler interface {
	TableName() string
}
