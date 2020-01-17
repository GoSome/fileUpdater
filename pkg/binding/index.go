/*
@Description: just go
@Author: skipper
@Date: 2020/1/15
@Time: 6:41 PM
@ProjectName fileUpdater
*/
package binding

import (
	"bytes"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"io"
)

func Index(c *gin.Context)  {
	template := rice.MustFindBox("dist")
	body,err := template.Bytes("index.html")
	if err != nil {
		c.String(400,"something goes wrong %s",err.Error())
		return
	}
	c.Header("Content-Type","text/html; charset=utf-8")
	_,err =io.Copy(c.Writer,bytes.NewReader(body))
	if err != nil {
		c.String(400,"something goes wrong %s",err.Error())
		return
	}
	c.Status(200)
}
