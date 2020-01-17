package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type UpdateContentRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func UpdateFile(c *gin.Context) {
	var req UpdateContentRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.String(400, "Unprocessable Data")
		return
	}

	updaters := Configs.GetUpdaterByName(req.Name)
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
