/*
@Description: just go
@Author: skipper
@Date: 2020/1/13
@Time: 5:13 PM
@ProjectName fileUpdater
*/
package server

import (
	"flag"
	"fmt"
	"github.com/GoSome/fileUpdater/pkg/config"
	"github.com/GoSome/fileUpdater/pkg/sig"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func Run() {
	fmt.Println("current pid: ", os.Getpid())
	sig.ListenSIGUSR2()
	go config.Watch()

	flag.StringVar(&config.Path, "config", "config.json", "server config file path")
	flag.Parse()

	config.Load()

	fmt.Println("configs:", config.Configs)
	app := gin.Default()
	app.GET("/api/updaters", GetUpdaters)
	app.GET("/api/updater", GetUpdater)
	app.GET("/api/content", GetContent)
	app.POST("/api/content", UpdateFile)

	log.Fatal(app.Run(config.Configs.ServerHost + ":" + config.Configs.ServerPort))
}
