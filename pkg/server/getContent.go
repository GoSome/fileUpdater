/*
@Description: just go
@Author: skipper
@Date: 2020/1/13
@Time: 5:08 PM
@ProjectName fileUpdater
*/
package server

import (
	"github.com/gin-gonic/gin"
)

func GetContent(c *gin.Context) {

	name := c.Query("name")
	if name == "" {
		// todo
		c.String(200, "nothing i found")
		return
	}
	u := Configs.GetUpdaterByName(name)
	if u == nil {
		c.String(400, "no idea")
		return
	}
	r, err := u.GetFile()
	defer r.Close()
	if err != nil {
		c.String(400, "no idea")
		return
	}
	// TODO: javascript treats json file content as object incorrectly
	content, err := u.GetFileContentAsString()
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(200, map[string]string{
		"content": content,
	})

}

type Req struct {
	Name string `json:"name"`
}
