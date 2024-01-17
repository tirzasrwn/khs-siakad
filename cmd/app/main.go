package main

import (
	"flag"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/tirzasrwn/khs-siakad/cmd/internal"
)

var (
	app internal.AppConfig
	il  *log.Logger
	el  *log.Logger
)

func main() {
	var loop bool
	flag.BoolVar(&loop, "l", false, "run in loop")
	flag.BoolVar(&loop, "loop", false, "run in loop")
	flag.Parse()

	il = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = il
	il = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)
	app.ErrorLog = el
	initializeAppConfig()
	s := internal.NewSiakad(app.SiakadUsername, app.SiakadPassword)
	if loop {
		loopRun(s)
	}
	oneRun(s)
}

func initializeAppConfig() {
	viper.SetConfigName("config.env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	app.SiakadPassword = viper.GetString("SIAKAD_PASSWORD")
	app.SiakadUsername = viper.GetString("SIAKAD_USERNAME")
	app.Webhook.URL = viper.GetString("WEBHOOK_URL")

	app.InfoLog.Println("success load config ...")
}
