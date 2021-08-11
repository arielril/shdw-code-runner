package router

import (
	"github.com/arielril/attack-graph-flow-code-runner/api/http/nmapapi"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	rt := gin.Default()

	rt.POST("/v1/nmap", nmapapi.ExecuteNmap)

	return rt
}
