package main

import (
	"github.com/drmendoz/backend-gin-gorm/routes"
	"github.com/drmendoz/backend-gin-gorm/utils"
)

func main() {
	port := ":" + utils.Viper.GetString("APP_PORT")
	_ = routes.R.Run(port)

}
