package sender

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"monitoring-system/scripts/agent/model"
)

type Sender struct {
	client *http.Client
	url    string
}

func NewSender(url string) *Sender {
	return &Sender{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
		url: url,
	}
}

func (s *Sender) Send(metric model.Metric) {
	jsonData, err := json.Marshal(metric)
	if err != nil {
		log.Println("erro ao converter JSON:", err)
		return
	}

	err = s.sendWithRetry(jsonData)
	if err != nil {
		log.Println("erro ao enviar metric:", err)
	} else {
		log.Printf("metric enviada: %+v\n", metric)
	}
}

func (s *Sender) sendWithRetry(data []byte) error {
	var err error

	for i := 0; i < 3; i++ {
		resp, err := s.client.Post(s.url, "application/json", bytes.NewBuffer(data))
		if err == nil {
			resp.Body.Close()
			return nil
		}

		log.Println("tentativa falhou, retry...", i+1)
		time.Sleep(2 * time.Second)
	}

	return err
}
