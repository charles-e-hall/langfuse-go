package langfuse

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/charles-e-hall/langfuse-go/transmission"
)

type Client struct {
	Url             string
	Credential      string
	MaxPendingItems int
	Transmission    transmission.Sender
}

type ClientOption func(*Client)

func WithCustomUrl(url string) ClientOption {
	return func(c *Client) {
		c.Url = url
	}
}

func WithMaxPendingItems(n int) ClientOption {
	return func(c *Client) {
		c.MaxPendingItems = n
	}
}

func GetCredentail() string {
	cred := base64.StdEncoding.EncodeToString(
		[]byte(
			fmt.Sprintf(
				"%s:%s",
				os.Getenv("LANGFUSE_PUBLIC_KEY"),
				os.Getenv("LANGFUSE_SECRET_KEY"),
			),
		),
	)
	return cred
}

func NewClient(opts ...ClientOption) (*Client, error) {
	c := &Client{}
	c.Transmission = &transmission.DefaultSender{}
	err := c.Transmission.Start()
	if err != nil {
		return c, err
	}
	c.Credential = GetCredentail()
	c.Transmission = transmission.NewDefaultSender(c.Url, c.Credential)

	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func (c *Client) Start() error {
	err := c.Transmission.Start()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Add(t transmission.Trace) error {
	err := c.Transmission.Add(t)
	if err != nil {
		return err
	}
	return nil
}
