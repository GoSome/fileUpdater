/*
@Description: just go
@Author: skipper
@Date: 2020/1/13
@Time: 5:13 PM
@ProjectName fileUpdater
*/
package server

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/GoSome/fileUpdater/pkg/binding"
	"github.com/GoSome/fileUpdater/pkg/config"
	"github.com/GoSome/fileUpdater/pkg/core"
	"github.com/gin-gonic/gin"
	"log"
)


func Run(cfg core.ServerConfigs) {
	app := gin.Default()
	app.Use(config.Inject)

	app.StaticFS("statics/", rice.MustFindBox("dist").HTTPBox())
	app.GET("/api/updaters", GetUpdaters)
	app.GET("/api/updater", GetUpdater)
	app.GET("/api/content", GetContent)
	app.POST("/api/content", UpdateFile)
	app.NoRoute(binding.Index)
	log.Fatal(app.Run(cfg.ServerHost + ":" + cfg.ServerPort))
}
