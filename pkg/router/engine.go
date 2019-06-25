package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iceyang/gin-exam/pkg/ctrl"
)

func Default() *gin.Engine {
	engine := gin.New()
	engine.Use(recovery)
	engine.Use(responseHandler)

	engine.GET("/ping", ctrl.Example.Ping)
	engine.GET("/404", ctrl.Example.NotFound)
	engine.GET("/ok", ctrl.Example.OK)
	return engine
}
