package queue

import (
	"github.com/nats-io/nats.go"
	"github.com/pableeee/processor/pkg/rest"
)

type NatsPublisher struct {
	conn *nats.Conn
}

func NewNatsPublisher() Publisher {
	return &NatsPublisher{}
}

func (nw *NatsPublisher) Publish(topic string, p []byte) error {
	if !nw.conn.IsConnected() {
		return ErrorConnection
	}

	err := nw.conn.Publish(topic, p)
	if err != nil {
		return err
	}

	nw.conn.Flush()

	return nil
}

func (nw *NatsPublisher) Close() {
	if !nw.conn.IsClosed() {
		nw.conn.Close()
	}
}

type NatsConsumer struct {
	conn *nats.Conn
}

func (nw *NatsPublisher) Subscribe(topic string, p Pusher) error {
	_, err := nw.conn.Subscribe(topic, func(msg *nats.Msg) {
		er := p.Push(msg.Data)
		if er != nil {
			// si falla el procesamiento, vuelvo a reencolar
			go func() {
				_ = nw.conn.Publish(topic, msg.Data)
				// TODO que pasa si no puedo reencolar??
			}()
		}
	})
	if err != nil {
		return err
	}

	return nil
}

func (nw *NatsConsumer) Close() {
	if !nw.conn.IsClosed() {
		nw.conn.Close()
	}
}

type HttpPusher struct {
	client *rest.Client
	url    string
}

func NewHttpPusher(url string) Pusher {
	p := HttpPusher{}
	p.url = url
	p.client = rest.NewRestClient()

	return &p
}

func (p *HttpPusher) Push(b []byte) error {

}
