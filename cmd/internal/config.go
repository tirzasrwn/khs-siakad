package internal

import (
	"log"

	"github.com/tirzasrwn/khs-siakad/webhook"
)

type AppConfig struct {
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	SiakadUsername string
	SiakadPassword string
	SiakadSemester string
	Webhook        webhook.Webhook
}
