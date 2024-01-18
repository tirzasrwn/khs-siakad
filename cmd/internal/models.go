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

func AreKHSEqual(slice1, slice2 []KHS) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
