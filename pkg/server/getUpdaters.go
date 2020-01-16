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
	"github.com/GoSome/fileUpdater/pkg/config"
	"github.com/gin-gonic/gin"
	"log"
)

func GetUpdaters(c *gin.Context) {
	updates := config.Configs.FileUpdaters
	log.Printf("updates: %s", updates)
	c.Header("Content-Type", "application/json")
	err := json.NewEncoder(c.Writer).Encode(&updates)
	if err != nil {
		log.Println("err: ", err.Error())
	}
	return
}
