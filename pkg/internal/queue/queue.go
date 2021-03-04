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
}
