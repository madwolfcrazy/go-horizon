package main

import (
	"fmt"
	"ymz465/go-horizon/config"
	"ymz465/go-horizon/model"
	"ymz465/go-horizon/router"

	"github.com/spf13/viper"
)

func main() {
	defaultPort := ":8080"
	config.Init()
	port := viper.GetString("server.port")
	if port == "" {
		port = defaultPort
	}
	//
	mysqlDSN := viper.GetString("db.mysqlDSN")
	model.InitDatabase(mysqlDSN)
	//
	r := router.InitRouter()
	fmt.Println("Now run service on: ", port)
	r.Run(port)
}
