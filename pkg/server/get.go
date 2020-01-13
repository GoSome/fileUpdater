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
	var gcr GetCReq
	err := c.Bind(&gcr)
	if err != nil {
		return
	}
	for _,u := range Configs.FileUpdaters {
		if u.Name == gcr.Name {
			r,err := u.GetFileContent()
			if err != nil {
				return
			}
			defer r.Close()
			io.Copy(c.Writer,r)
			c.Status(200)
			return
		}
	}
	c.String(400,"no idea!")
}

type GetCReq struct {
	Name string `json:"name"`

}