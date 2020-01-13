/*
@Description: just go
@Author: skipper
@Date: 2020/1/13
@Time: 5:12 PM
@ProjectName fileUpdater
*/
package server

import "github.com/gin-gonic/gin"

func GetUpdaters(c *gin.Context) {
	updates := Configs.FileUpdaters
	c.JSON(200, updates)
	return
}
