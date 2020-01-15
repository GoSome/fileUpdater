package sig

import (
	"fmt"
	"github.com/GoSome/fileUpdater/pkg/config"
	"os"
	"os/signal"
	"syscall"
)

func ListenSIGUSR2() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR2)

	go func() {
		for {
			<-s
			config.Load()
			fmt.Println("config reloaded")
		}
	}()
}
