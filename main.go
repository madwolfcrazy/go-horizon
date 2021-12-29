package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"ymz465/go-horizon/config"
	"ymz465/go-horizon/model"
	"ymz465/go-horizon/router"

	"github.com/madwolfcrazy/endless"
	"github.com/spf13/viper"
)

func main() {
	config.Init()
	// is exit flag
	if len(os.Args) > 1 && os.Args[1] == "stop" {
		// kill pid
		cmdStr := fmt.Sprintf("kill -EXIT `cat %s`", config.PIDFile)
		exec.Command(cmdStr)
		os.Exit(0)
	}
	// is restart flag
	if len(os.Args) > 1 && os.Args[1] == "restart" {
		// send HUP singal
		cmdStr := fmt.Sprintf("kill -HUP `cat %s`", config.PIDFile)
		exec.Command(cmdStr)
		os.Exit(0)
	}
	var err error
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
	/*
		generater tls key pair
		# openssl genrsa -out ./server.key 2048
		# openssl req -new -x509 -key ./server.key -out ./server.pem -days 365
	*/
	//err = server.ListenAndServeTLS(config.TLSPublicKey, config.TLSPrivateKey)
	if err != nil {
		log.Println(err)
	}

}
