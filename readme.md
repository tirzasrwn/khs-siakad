# khs-siakad

## About

Scraper for https://siakad.utdi.ac.id/ to check newst KHS.
The issue with Siakad is that there is no notification feature when the grades are input by the lecturer.
So the idea is using scrapper to check if data changed or not, then notify it.

## Requirement

- go
- google-chrome or chromium
- discord webhook url (if you want to notify)

# Stack

- go
- chromedp

## Build and Run

```sh
# copy config.env.example to config.env
# change username and password in config.env

# build
go build -o khs ./cmd/app

# run in one mode
./khs
# run in loop mode, it will notify you if the data changed
# notofication is send by discord webhook
./khs -l
```

## Todo

- [ ] demo vidio
