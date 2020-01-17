/*
@Description: just go
@Author: skipper
@Date: 2020/1/13
@Time: 5:13 PM
@ProjectName fileUpdater
*/
package server

import (
	"encoding/json"
	"flag"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/GoSome/fileUpdater/pkg/binding"
	"github.com/GoSome/fileUpdater/pkg/types"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

var configPath string
var Configs types.ServerConfigs

func Run() {
	flag.StringVar(&configPath, "config", "config.json", "server config file path")
	flag.Parse()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file \"%s\" not exist", configPath)
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("open config file \"%s\" failed", configPath)
	}

	if strings.HasSuffix(configPath, ".json") {
		//
		err = json.NewDecoder(configFile).Decode(&Configs)
		if err != nil {
			// todo
			panic(err)
		}
	} else if strings.HasSuffix(configPath, ".yaml") {
		log.Println("what")
		err := yaml.NewDecoder(configFile).Decode(&Configs)
		if err != nil {
			// todo
			log.Println("err: ", err.Error())
			panic(err)
		}
	} else {
		panic("config file path must end with .json or .yaml")
	}

	fmt.Println("configs:", Configs)
	app := gin.Default()

	app.StaticFS("statics/",rice.MustFindBox("dist").HTTPBox())
	app.GET("/api/updaters", GetUpdaters)
	app.GET("/api/updater", GetUpdater)
	app.GET("/api/content", GetContent)
	app.POST("/api/content", UpdateFile)
	app.NoRoute(binding.Index)
	log.Fatal(app.Run(Configs.ServerHost + ":" + Configs.ServerPort))
}
