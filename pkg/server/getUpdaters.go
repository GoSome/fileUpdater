/*
@Description: just go
@Author: skipper
@Date: 2020/1/13
@Time: 5:12 PM
@ProjectName fileUpdater
*/
package server

import (
	"encoding/json"
	"log"

	"github.com/GoSome/fileUpdater/pkg/config"
	"github.com/gin-gonic/gin"
)

func (a *App) GetUpdaters(c *gin.Context) {
	updates := config.Config.FileUpdaters
	log.Printf("updates: %v", updates)
	c.Header("Content-Type", "application/json")
	err := json.NewEncoder(c.Writer).Encode(&updates)
	if err != nil {
		log.Println("err: ", err.Error())
	}
	return
}
