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
	"github.com/GoSome/fileUpdater/pkg/core"
	"github.com/gin-gonic/gin"
)

type getUpdaterResponse struct {
	Updater *core.FileUpdater `json:"updater"`
	Content string            `json:"content"`
}

func GetUpdater(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.String(404, "Not Found")
		return
	}

	getConfig, _ := c.Get("cfg")
	cfg := getConfig.(core.ServerConfigs)

	u := cfg.GetUpdaterByName(name)
	if u == nil {
		c.String(404, "Not Found")
		return
	}

	content, err := u.GetFileContentAsString()
	if err != nil {
		c.String(400, err.Error())
		return
	}

	response := getUpdaterResponse{
		Updater: u,
		Content: content,
	}

	if err := json.NewEncoder(c.Writer).Encode(&response); err != nil {
		c.String(500, "Unknown error: %s", err)
	}
}
