package queue

import (
	"github.com/nats-io/nats.go"
)

type NatsWriter struct {
	conn *nats.Conn
}

func (nw *NatsWriter) Publish(topic string, p []byte) (int, error) {
	if !nw.conn.IsConnected() {
		return 0, ErrorConnection
	}

	err := nw.conn.Publish(topic, p)
	if err != nil {
		return 0, err
	}

	nw.conn.Flush()

	return len(p), nil
}

func (nw *NatsWriter) Subscribe(topic string, cb func(b []byte)) error {
	_, err := nw.conn.Subscribe(topic, func(msg *nats.Msg) {
		cb(msg.Data)
	})
	if err != nil {
		return err
	}

	return nil
}

func (nw *NatsWriter) Close() {
	if !nw.conn.IsClosed() {
		nw.conn.Close()
	}
}
