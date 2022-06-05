package server

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/GoSome/fileUpdater/pkg/core"
	"github.com/gin-gonic/gin"
)

type UpdateContentRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (a *App) UpdateFile(c *gin.Context) {
	var req UpdateContentRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.String(400, "Unprocessable Data")
		return
	}

	getConfig, _ := c.Get("cfg")
	cfg := getConfig.(core.ServerConfigs)

	updaters := cfg.GetUpdaterByName(req.Name)
	if updaters == nil {
		c.String(404, "Not Found")
		return
	}

	// TODO: get file content from FormFile
	f := strings.NewReader(req.Content)
	if err := updaters.UpdateFile(f); err != nil {
		//todo
		log.Printf("something bad when update file %s", err)
		c.String(400, "%s", err.Error())
		return
	}
	c.String(200, "all is well")
	return
}
