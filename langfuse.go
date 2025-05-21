package langfuse

import (
	"os"

	"github.com/charles-e-hall/langfuse-go/transmission"
)

type Client struct {
	Url          string
	PublicKey    string
	SecretKey    string
	Transmission transmission.Sender
}

type ClientOption func(*Client)

func WithCustomUrl(url string) ClientOption {
	return func(c *Client) {
		c.Url = url
	}
}

func NewClient(opts ...ClientOption) (*Client, error) {
	c := &Client{}
	c.Transmission = &transmission.DefaultSender{}
	err := c.Transmission.Start()
	if err != nil {
		return c, err
	}
	c.PublicKey = os.Getenv("LANGFUSE_PUBLIC_KEY")
	c.SecretKey = os.Getenv("LANGFUSE_SECRET_KEY")

	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}
