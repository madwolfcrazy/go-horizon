package main

import (
	"fmt"
	"log"
	"os"
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
	server := endless.NewServer(port, r)
	server.BeforeBegin = func(add string) {
		pid := syscall.Getpid()
		log.Printf("Now service runing on: %s, and pid is %d", port, pid)
		// save pid if need
		pidFile, err := os.Create("./logs/run.pid")
		if err == nil {
			pidFile.WriteString(fmt.Sprint(pid))
			pidFile.Close()
		}
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
