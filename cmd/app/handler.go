package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/tirzasrwn/khs-siakad/cmd/internal"
	"github.com/tirzasrwn/khs-siakad/webhook"
)

func oneRun(s *internal.Siakad) {
	app.InfoLog.Println(s.GetKHSData())
}

func loopRun(s *internal.Siakad) {
	var currentKHS []internal.KHS
	w := webhook.New(app.Webhook.URL)
	for {
		newKHS, err := s.GetKHSData()
		if err != nil {
			app.ErrorLog.Println(err)
		}
		app.InfoLog.Println(newKHS)
		if !reflect.DeepEqual(currentKHS, newKHS) {
			currentKHS = newKHS
			body, err := w.SendMessage(fmt.Sprint(currentKHS))
			if err != nil {
				app.ErrorLog.Printf("error send message: %s got return: %s", err, body)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
