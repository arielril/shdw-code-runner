package router

import (
	"github.com/arielril/attack-graph-flow-code-runner/internal/nmap"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	rt := gin.Default()

	rt.POST("/v1/nmap", nmap.ExecuteNmap)

	return rt
}
