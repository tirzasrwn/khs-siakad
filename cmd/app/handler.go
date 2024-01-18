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
			time.Sleep(10 * time.Second)
			continue
		}
		app.InfoLog.Println(newKHS)
		if !reflect.DeepEqual(currentKHS, newKHS) {
			currentKHS = newKHS
			loc, _ := time.LoadLocation("Asia/Jakarta")
			body, err := w.SendMessage(fmt.Sprintf("INFO SIAKAD %s %s", time.Now().In(loc).Format(time.RFC3339), currentKHS))
			if err != nil {
				app.ErrorLog.Printf("error send message: %s got return: %s", err, body)
				time.Sleep(10 * time.Second)
				continue
			}
		}
		time.Sleep(10 * time.Second)
	}
}
