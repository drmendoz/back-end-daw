package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/drmendoz/backend-gin-gorm/utils"
)

var Db *gorm.DB

func init() {
	user := utils.Viper.GetString("DB_USER")
	password := utils.Viper.GetString("DB_PASS")
	server := utils.Viper.GetString("DB_SERVER")
	database := utils.Viper.GetString("DB_NAME")
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
	err := Db.AutoMigrate(&Usuario{}, &Administrador{}, &Cliente{}, &Producto{}, &Pedido{}, &DetallePedido{})
	if err != nil {
		utils.Log.Warn(err)
		utils.Log.Fatal("Error al migrar modelos")

	}
}

type Tabler interface {
	TableName() string
}
