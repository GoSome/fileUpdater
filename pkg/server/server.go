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
	"github.com/GoSome/fileUpdater/pkg/types"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var configPath string
var Configs types.ServerConfigs

func Run() {
	flag.StringVar(&configPath, "config", "config.json", "server config file path")
	flag.Parse()

	configFile, err := os.Open(configPath)
	if err != nil {
		// todo
		panic(err)
		return
	}
	err = json.NewDecoder(configFile).Decode(&Configs)
	if err != nil {
		// todo
		panic(err)
		return
	}
	fmt.Println(Configs)
	app := gin.Default()
	app.GET("/updates", GetUpdaters)
	app.POST("/content",GetContent)

	log.Fatal(app.Run(Configs.ServerHost + ":" + Configs.ServerPort))
}
