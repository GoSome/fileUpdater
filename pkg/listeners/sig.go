package listeners

import (
	"github.com/GoSome/fileUpdater/pkg/config"
	"log"
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
			log.Println("config file has been changed, attempt to reload...")
			config.Parse(false)
		}
	}()
}
