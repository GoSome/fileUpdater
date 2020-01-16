/*
@Description: just go
@Author: skipper
@Date: 2020/1/13
@Time: 5:08 PM
@ProjectName fileUpdater
*/
package server

import (
	"github.com/GoSome/fileUpdater/pkg/config"
	"github.com/gin-gonic/gin"
	"io"
)

func GetContent(c *gin.Context) {

	name := c.Query("name")
	if name == "" {
		// todo
		c.String(200, "nothing i found")
		return
	}
	u := config.Configs.GetUpdaterByName(name)
	if u == nil {
		c.String(400, "no idea")
		return
	}
	r, err := u.GetFileContent()
	defer r.Close()
	if err != nil {
		c.String(400, "no idea")
		return
	}
	io.Copy(c.Writer, r)
	c.Status(200)

}

type Req struct {
	Name string `json:"name"`
}
