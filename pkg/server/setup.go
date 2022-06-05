package server

import (
	"github.com/GoSome/fileUpdater/pkg/core"
	"github.com/gin-gonic/gin"
)

type App struct {
	Options core.ServerConfigs
	Engine  *gin.Engine
}
