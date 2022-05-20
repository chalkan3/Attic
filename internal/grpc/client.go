package grpc

import (
	connector "google.golang.org/grpc"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

func (c *Client) Dial() (*connector.ClientConn, error) {
	return connector.Dial(c.addr, connector.WithInsecure())

}
