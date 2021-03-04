package queue

import (
	"fmt"
)

var ErrorConnection = fmt.Errorf("not connected")

type Queue interface {
	Write(p []byte) (int, error)
	Subscribe(topic string, cb func(b []byte)) error
	Close()
}

type Publisher interface {
	Publish(topic string, b []byte) error
	Close()
}

type Pusher interface {
	Push(b []byte) error
}

type Consumer interface {
	Subscribe(topic string, p Pusher) error
	Close()
}
