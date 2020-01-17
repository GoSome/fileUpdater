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
	"github.com/GoSome/fileUpdater/pkg/server"
	"github.com/sevlyar/go-daemon"
	"log"
)

func main() {
	cntxt := &daemon.Context{
		PidFileName: "sample.pid",
		PidFilePerm: 0644,
		LogFileName: "sample.log",
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
	server.Run()
}
