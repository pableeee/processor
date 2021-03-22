package osutil

import (
	"os"
	"os/signal"
)

func ExecOn(f func(), sig ...os.Signal) {
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, sig...)
		<-sc
		f()
	}()
}
