package transmission

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

type Sender interface {
	SendAll() error
	Start() error
	Worker() error
	Flush() error
	Add(t Trace) error
	ReadAllJSON() string
}

type DefaultSender struct {
	SendQueue       []Trace
	Started         bool
	MaxPendingItems int
	Url             string
	Credential      string
	lock            sync.Mutex
}

func (s *DefaultSender) SendAll() error {
	body := bytes.NewBuffer([]byte("this stuff"))
	req, err := http.NewRequest("POST", s.Url, body)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", s.Credential))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error msg: %s", resp.Body)
	}
	return nil
}

func (s *DefaultSender) Start() error {
	s.Started = true
	go s.Worker()
	return nil
}

func (s *DefaultSender) Worker() error {
	for s.Started == true {
		if len(s.SendQueue) >= s.MaxPendingItems {
			err := s.SendAll()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *DefaultSender) Flush() error {
	err := s.SendAll()
	if err != nil {
		return err
	}
	return nil
}

func (s *DefaultSender) Add(t Trace) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.SendQueue = append(s.SendQueue, t)
	return nil
}

func (s *DefaultSender) ReadAllJSON() string {
	return ""
}

func NewDefaultSender(url string, credential string) *DefaultSender {
	s := DefaultSender{
		Url:             url,
		MaxPendingItems: 5,
		Credential:      credential,
	}
	s.SendQueue = make([]Trace, s.MaxPendingItems)
	s.Started = false
	return &s
}
