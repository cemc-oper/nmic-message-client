package sender

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type HttpSender struct {
	PostUrl        string
	RequestTimeout time.Duration
}

func (s *HttpSender) SendMessage(blob []byte) error {
	client := resty.New()
	client.SetTimeout(s.RequestTimeout)
	client.SetDebug(true)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(blob).
		Post(s.PostUrl)

	if err != nil {
		return err
	}

	fmt.Printf("%v\n", resp)

	return nil
}
