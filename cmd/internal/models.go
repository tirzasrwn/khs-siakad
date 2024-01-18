package internal

import "context"

type Siakad struct {
	username       string
	password       string
	chromedpCtx    context.Context
	ChromedpCancel context.CancelFunc
}

type KHS struct {
	MataKuliah string
	Nilai      string
}
