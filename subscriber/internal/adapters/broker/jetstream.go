package broker

import (
	"fmt"

	"github.com/nats-io/nats.go"

	"subscriber/pkg/config"
)


type NatsConn struct{
	Con *nats.Conn
	cfg *config.JetConfig
}

func New(cfg *config.JetConfig) (*NatsConn, error) {
	ncon, err := nats.Connect(cfg.Address)
	if err != nil {
		return nil, fmt.Errorf("nats conn: %w", err)
	} 
	return &NatsConn{
		Con: ncon,
		cfg: cfg,
	}, nil
}


func (nc *NatsConn) Close(){
	nc.Con.Close()
}


