/*
@Description: just go
@Author: skipper
@Date: 2020/1/13
@Time: 5:13 PM
@ProjectName fileUpdater
*/
package server

import (
	"log"

	//	rice "github.com/GeertJohan/go.rice"
	"github.com/GoSome/fileUpdater/pkg/binding"
	"github.com/GoSome/fileUpdater/pkg/config"
	"github.com/GoSome/fileUpdater/pkg/core"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(cfg core.ServerConfigs) {
	cfg.RunProcess()
	app := gin.Default()
	app.Use(config.Inject)
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
	app.GET("/api/updaters", GetUpdaters)
	app.GET("/api/updater", GetUpdater)
	app.GET("/api/content", GetContent)
	app.POST("/api/content", UpdateFile)
	app.POST("/api/exec", Exec)
	if !cfg.DisableUI {
		//		app.StaticFS("statics/", rice.MustFindBox("dist").HTTPBox())
		app.NoRoute(binding.Index)
	}
	app.NoRoute(GetUpdaters)
	log.Fatal(app.Run(cfg.ServerHost + ":" + cfg.ServerPort))
}
