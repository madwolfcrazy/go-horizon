package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"ymz465/go-horizon/config"
	"ymz465/go-horizon/llog"
	"ymz465/go-horizon/model"
	"ymz465/go-horizon/router"

	"github.com/madwolfcrazy/endless"
	"github.com/spf13/viper"
)

func main() {
	config.Init()
	var err error
	if viper.GetString("runmode") != "debug" {
		//
		logFile, err := llog.GetLoggerFile(config.RunLog)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer logFile.Close()
		log.SetOutput(logFile)
	}
	//
	defaultPort := ":8080"
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
		pidFile, err := os.Create(config.PIDFile)
		if err == nil {
			pidFile.WriteString(fmt.Sprint(pid))
			pidFile.Close()
		}
	}
	log.Println("start go-horizon")
	log.Printf("%s - %d ", "start go-horizon", 1978)
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
