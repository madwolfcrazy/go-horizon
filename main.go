package main

import (
	"log"
	"syscall"
	"ymz465/go-horizon/config"
	"ymz465/go-horizon/model"
	"ymz465/go-horizon/router"

	"github.com/fvbock/endless"
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
	//r.Run(port)
	server := endless.NewServer(port, r)
	server.BeforeBegin = func(add string) {
		log.Printf("Now service runing on: %s, and pid is %d", port, syscall.Getpid())
		// save pid if need
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
