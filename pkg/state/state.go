package state

import (
	"context"
	"fmt"
)

type stateMachine struct {
	state   string
	actionc chan func()
}

func newStateMachine() *stateMachine {
	return &stateMachine{
		state:   "initial",
		actionc: make(chan func()),
	}
}

func (s *stateMachine) Run(ctx context.Context) error {
	for {
		select {
		case f := <-s.actionc:
			f()
		case <-ctx.Done():
			return fmt.Errorf("context cancelled: %w", ctx.Err())
		}
	}
}

func (s *stateMachine) foo() int {
	c := make(chan int)
	s.actionc <- func () {
		if s.state == "a" {
			s.state = "b"
		}
		c <- 123
	}

	return <- c
}