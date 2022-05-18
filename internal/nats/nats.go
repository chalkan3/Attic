package nats_cli

import "github.com/nats-io/nats.go"

type NATS struct {
	Conn *nats.Conn
}

func (nats *NATS) GetConn() *nats.Conn {
	return nats.Conn
}

func NewNATS(uri string, defaults bool) (*NATS, error) {
	if defaults {
		uri = nats.DefaultURL
	}

	nc, err := nats.Connect(uri)
	if err != nil {
		return nil, err
	}

	return &NATS{
		Conn: nc,
	}, nil

}
