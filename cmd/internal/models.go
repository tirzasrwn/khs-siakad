package internal

import (
	"context"
	"fmt"
	"strings"
)

type Siakad struct {
	username       string
	password       string
	semester       string
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

func KHSPrint(khs []KHS) {
	// Calculate the maximum length of the course names and grades for formatting
	maxMataKuliahLength := 0
	maxNilaiLength := 0
	for _, kh := range khs {
		if len(kh.MataKuliah) > maxMataKuliahLength {
			maxMataKuliahLength = len(kh.MataKuliah)
		}
		if len(kh.Nilai) > maxNilaiLength {
			maxNilaiLength = len(kh.Nilai)
		}
	}

	// Define the number of columns for the table
	columns := 1
	rows := (len(khs) + columns - 1) / columns

	// Print the courses and grades in table format
	for r := 0; r < rows; r++ {
		rowStr := ""
		for c := 0; c < columns; c++ {
			index := r + c*rows
			if index < len(khs) {
				rowStr += fmt.Sprintf("%-*s %-*s ", maxMataKuliahLength, khs[index].MataKuliah, maxNilaiLength, khs[index].Nilai)
			}
		}
		fmt.Println(strings.TrimSpace(rowStr))
	}
}
