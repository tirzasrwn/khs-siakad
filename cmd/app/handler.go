package main

import (
	"fmt"
	"os"
	"os/exec"
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
			app.InfoLog.Println("--DATA CHANGED--")
			cmd := exec.Command("dunstify", "-u", "normal", "-t", "0",
				fmt.Sprintf("\"%s %s\"\n", "--Data Changed--", fmt.Sprint(currentKHS)))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				app.ErrorLog.Println(err)
			}
			w.SendMessage(fmt.Sprint(currentKHS))
		}
		time.Sleep(10 * time.Second)
	}
}
