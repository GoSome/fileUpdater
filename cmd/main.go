/*
@Description: just go
@Author: skipper
@Date: 2020/1/13
@Time: 5:49 PM
@ProjectName fileUpdater
*/
package main

import (
	"flag"

	"github.com/GoSome/fileUpdater/pkg/config"
	"github.com/GoSome/fileUpdater/pkg/listeners"

	"github.com/GoSome/fileUpdater/pkg/server"

	"log"

	"github.com/sevlyar/go-daemon"
)

func main() {
	flag.StringVar(&config.Path, "config", "config.json", "server config file path")
	flag.BoolVar(&config.DaemonZ, "d", false, "daemon")
	flag.BoolVar(&config.IncludeSelf, "i", false, "include config file to updaters")
	flag.BoolVar(&config.DisableHotReload, "disable-reload", false, "disable hot reload config file")
	flag.StringVar(&config.PidPath, "pid", "", "pid path work in daemon")
	flag.StringVar(&config.LogFile, "log", "", "log path work in daemon")
	flag.Parse()

	config.Parse(true)

	if !config.DisableHotReload {
		listeners.ListenSIGUSR2()
		go config.Watch()
	}

	if config.DaemonZ {
		cntxt := &daemon.Context{
			PidFileName: config.PidPath,
			PidFilePerm: 0644,
			LogFileName: config.LogFile,
			LogFilePerm: 0640,
			WorkDir:     "./",
			Umask:       027,
			Args:        flag.Args(),
		}

		d, err := cntxt.Reborn()
		if err != nil {
			log.Fatal("Unable to run: ", err)
		}
		if d != nil {
			return
		}
		defer cntxt.Release()

		log.Print("- - - - - - - - - - - - - - -")
		log.Print("daemon started")
	}
	server.Run(config.Config)
}
