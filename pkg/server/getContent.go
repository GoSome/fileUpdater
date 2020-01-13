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
	"io"
)

func GetContent(c *gin.Context) {
	var gcr Req
	err := c.Bind(&gcr)
	if err != nil {
		return
	}
	u := Configs.GetUpdatersByName(gcr.Name)
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
