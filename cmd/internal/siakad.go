package internal

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

var s *Siakad

func NewSiakad(username string, password string, semester string) (*Siakad, error) {
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())

	// Start the browser without timeout
	err := chromedp.Run(ctx)
	if err != nil {
		return nil, err
	}

	s = &Siakad{
		username:       username,
		password:       password,
		semester:       semester,
		chromedpCtx:    ctx,
		ChromedpCancel: cancel,
	}

	return s, nil
}

func (s *Siakad) GetKHSData() ([]KHS, error) {
	// Create timeout context
	ctx, cancel := context.WithTimeout(s.chromedpCtx, 60*time.Second)
	defer cancel()

	// Open new tab because browser already started in NewSiakad
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// Navigate to the login page
	err := chromedp.Run(ctx, chromedp.Navigate(baseURL))
	if err != nil {
		return nil, err
	}
	// Wait for the login page to load
	err = chromedp.Run(ctx, chromedp.WaitVisible("#username"))
	if err != nil {
		return nil, err
	}
	// Fill in the username and password fields
	err = chromedp.Run(ctx,
		chromedp.SendKeys("#username", s.username),
		chromedp.SendKeys("#password", s.password),
	)
	if err != nil {
		return nil, err
	}

	// Click the login button
	err = chromedp.Run(ctx, chromedp.Click(".button[type='submit']"))
	if err != nil {
		return nil, err
	}
	// Wait for the successful login page to load
	err = chromedp.Run(ctx, chromedp.WaitVisible("#user-info"))
	if err != nil {
		return nil, err
	}

	// Click the link with the text "Kartu Rencana Studi"
	err = chromedp.Run(ctx, chromedp.Click("/html/body/div/div[2]/div[2]/div[2]/ul/li[6]"))
	if err != nil {
		return nil, err
	}

	// Wait for select semester
	err = chromedp.Run(ctx, chromedp.WaitVisible(`select[name="lstSemester"]`))
	if err != nil {
		return nil, err
	}

	// Choose semester by select value
	err = chromedp.Run(ctx,
		chromedp.SetValue(`select[name="lstSemester"]`, s.semester),
	)
	if err != nil {
		return nil, err
	}

	// Click lihat
	err = chromedp.Run(ctx,
		chromedp.Click("/html/body/div/div[2]/div[1]/form[1]/table/tbody/tr/td[2]/input"),
		chromedp.WaitReady(".table-common"),
	)
	if err != nil {
		return nil, err
	}

	// printBody(ctx)
	return parseKHS(filterData(getBody(ctx))), nil
}

func printBody(ctx context.Context) {
	fmt.Println(getBody(ctx))
}

func getBody(ctx context.Context) string {
	var bodyContent string
	chromedp.Run(ctx, chromedp.InnerHTML("body", &bodyContent))
	return bodyContent
}

func filterData(output string) string {
	// Define regular expression patterns
	pattern1 := `<td width="44%">([^<]+)</td>`
	pattern2 := `<td width="6%" align="center">[ABCD]?</td>`

	// Compile regular expressions
	re1 := regexp.MustCompile(pattern1)
	re2 := regexp.MustCompile(pattern2)

	// Split the output into lines

	lines := strings.Split(string(output), "\n")

	// Initialize a string to store the matches
	var resultString string

	// Process each line
	for _, line := range lines {
		// Match against pattern 1
		matches1 := re1.FindStringSubmatch(line)
		if len(matches1) > 1 {
			resultString += matches1[1] + "\n"
		}

		// Match against pattern 2
		matches2 := re2.FindStringSubmatch(line)
		if len(matches2) > 0 {
			openTag := `<td width="6%" align="center">`
			closeTag := `</td>`

			startIndex := strings.Index(matches2[0], openTag)
			endIndex := strings.Index(matches2[0], closeTag)

			if startIndex != -1 && endIndex != -1 {
				value := matches2[0][startIndex+len(openTag) : endIndex]
				resultString += value + "\n"
			}
		}
	}

	return resultString
}

func parseKHS(input string) []KHS {
	lines := strings.Split(input, "\n")
	var khsList []KHS

	for i := 0; i < len(lines)-1; i += 2 {
		khs := KHS{
			MataKuliah: lines[i],
			Nilai:      lines[i+1],
		}
		khsList = append(khsList, khs)
	}

	return khsList
}
