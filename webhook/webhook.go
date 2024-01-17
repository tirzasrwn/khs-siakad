package webhook

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Webhook struct {
	URL string
}

var w *Webhook

func New(url string) *Webhook {
	w = &Webhook{
		URL: url,
	}
	return w
}

func (w *Webhook) SendMessage(message string) (string, error) {
	method := "POST"
	payload := strings.NewReader(fmt.Sprintf(`{
    "content": "%s"
    }`, message))
	client := &http.Client{}
	req, err := http.NewRequest(method, w.URL, payload)
	if err != nil {
		return "", fmt.Errorf("can not create request: %s", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
