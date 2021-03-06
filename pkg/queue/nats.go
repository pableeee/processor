package queue

import (
	"bytes"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/pableeee/processor/pkg/rest"
)

type NatsPublisher struct {
	conn *nats.Conn
}

func NewNatsPublisher(address string, port int) (Publisher, error) {
	url := fmt.Sprintf("%s:%d", address, port)
	opts := []nats.Option{nats.Name("NATS Sample Publisher")}

	nc, err := nats.Connect(url, opts...)
	if err != nil {
		return nil, fmt.Errorf("error connecting to nats server %w", err)
	}

	return &NatsPublisher{conn: nc}, nil
}

func (nw *NatsPublisher) Publish(topic string, p []byte) error {
	if !nw.conn.IsConnected() {
		return ErrorConnection
	}

	err := nw.conn.Publish(topic, p)
	if err != nil {
		return fmt.Errorf("error publishing message %w", err)
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

func NewNatsConsumer(address string, port int) (Consumer, error) {
	url := fmt.Sprintf("%s:%d", address, port)
	opts := []nats.Option{nats.Name("NATS Sample Publisher")}

	nc, err := nats.Connect(url, opts...)
	if err != nil {
		return nil, fmt.Errorf("error connecting to nats server %w", err)
	}

	return &NatsPublisher{conn: nc}, nil
}

func (nw *NatsPublisher) Subscribe(topic string, p Pusher) error {
	_, err := nw.conn.Subscribe(topic, func(msg *nats.Msg) {
		er := p.Push(msg.Data)
		if er != nil {
			// si falla el procesamiento, vuelvo a reencolar
			go func() {
				_ = nw.conn.Publish(topic, msg.Data)
				// que pasa si no puedo reencolar??
				// o usar nates reply
			}()
		}
	})
	if err != nil {
		return fmt.Errorf("error subscribing to topic %w", err)
	}

	return nil
}

func (nw *NatsConsumer) Close() {
	if !nw.conn.IsClosed() {
		nw.conn.Close()
	}
}

type HTTPPusher struct {
	client  *rest.Client
	url     string
	headers map[string]string
}

func NewHTTPPusher(url string, headers map[string]string) Pusher {
	p := HTTPPusher{
		client:  rest.NewRestClient(),
		url:     url,
		headers: headers,
	}

	return &p
}

func (p *HTTPPusher) Close() {

}

func (p *HTTPPusher) Push(b []byte) error {
	_, err := p.client.Execute("POST", p.url, bytes.NewReader(b), p.headers)

	return fmt.Errorf("error pushing to http endponint %w", err)
}
