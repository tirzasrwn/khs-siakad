package internal

import (
	"log"
)

type AppConfig struct {
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	SiakadUsername string
	SiakadPassword string
}
