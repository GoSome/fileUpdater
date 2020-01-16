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
)

func GetUpdater(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.String(404, "Not Found")
		return
	}

	u := config.Configs.GetUpdaterByName(name)
	if u == nil {
		c.String(404, "Not Found")
		return
	}

	if err := json.NewEncoder(c.Writer).Encode(&u); err != nil {
		c.String(500, "Unknown error: %s", err)
	}
}
