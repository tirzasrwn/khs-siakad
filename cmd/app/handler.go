package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/tirzasrwn/khs-siakad/cmd/internal"
	"github.com/tirzasrwn/khs-siakad/webhook"
)

func oneRun(s *internal.Siakad) {
	khs, err := s.GetKHSData()
	if err != nil {
		app.ErrorLog.Println(err)
	}
	internal.KHSPrint(khs)
	fmt.Print("press enter to exit...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}

func loopRun(s *internal.Siakad) {
	var currentKHS []internal.KHS
	w := webhook.New(app.Webhook.URL)
	loc, _ := time.LoadLocation("Asia/Jakarta")
	for {
		newKHS, err := s.GetKHSData()
		if err != nil {
			app.ErrorLog.Println(err)
			time.Sleep(30 * time.Second)
			continue
		}
		if len(newKHS) == 0 {
			app.ErrorLog.Println(errors.New("get khs data returns nothing"))
			time.Sleep(30 * time.Second)
			continue
		}
		app.InfoLog.Println(newKHS)
		if !internal.AreKHSEqual(currentKHS, newKHS) {
			currentKHS = newKHS
			body, err := w.SendMessage(fmt.Sprintf("--> %s %s", time.Now().In(loc).Format(time.RFC3339), currentKHS))
			if err != nil {
				app.ErrorLog.Printf("error send message: %s got return: %s", err, body)
				time.Sleep(30 * time.Second)
				continue
			}
		}
		time.Sleep(30 * time.Second)
	}
}
