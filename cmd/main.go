/*
@Description: just go
@Author: skipper
@Date: 2020/1/13
@Time: 5:49 PM
@ProjectName fileUpdater
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/GoSome/fileUpdater/pkg/core"
	"github.com/GoSome/fileUpdater/pkg/server"
	"github.com/sevlyar/go-daemon"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

var configPath string
var daemonZ bool
var pidPath string
var logFile string
var config core.ServerConfigs

func main() {
	flag.StringVar(&configPath, "config", "config.json", "server config file path")
	flag.BoolVar(&daemonZ,"d",false,"daemon")
	flag.StringVar(&pidPath,"pid","","pid path work in daemon")
	flag.StringVar(&logFile,"log","","log path work in daemon")
	flag.Parse()

	ParseConfig()
	if daemonZ {
		cntxt := &daemon.Context{
			PidFileName: pidPath,
			PidFilePerm: 0644,
			LogFileName: logFile,
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

	server.Run(config)
}


func ParseConfig()  {

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file \"%s\" not exist", configPath)
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("open config file \"%s\" failed", configPath)
		return
	}
	if strings.HasSuffix(configPath, ".json") {
		//
		err = json.NewDecoder(configFile).Decode(&config)
		if err != nil {
			// todo
			panic(err)
		}
	} else if strings.HasSuffix(configPath, ".yaml") {
		err := yaml.NewDecoder(configFile).Decode(&config)
		if err != nil {
			// todo
			log.Println("err: ", err.Error())
			panic(err)
		}
	} else {
		panic("config file path must end with .json or .yaml")
	}

	configFile.Close()
}