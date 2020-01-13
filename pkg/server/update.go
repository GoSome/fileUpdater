package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

func UpdateFile(c *gin.Context) {
	name := c.PostForm("name")
	updaters := Configs.GetUpdatersByName(name)
	if updaters != nil {
		mpf, err := c.FormFile("file")
		if err != nil {
			//todo
			log.Println("something bad")
			c.String(400, "%s", err.Error())
			return
		}
		f, err := mpf.Open()
		err = updaters.UpdateFile(f)
		if err != nil {
			//todo
			log.Println("something bad when update file")
			c.String(400, "%s", err.Error())
			return
		}
		c.String(200, "all is well")
		return
	}
	c.String(200, "nothing updated")

}
