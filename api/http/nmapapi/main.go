package nmapapi

import (
	"net/http"

	"github.com/arielril/attack-graph-flow-code-runner/internal/command"
	"github.com/arielril/attack-graph-flow-code-runner/internal/nmap"
	"github.com/arielril/attack-graph-flow-code-runner/log"
	"github.com/gin-gonic/gin"
)

var logger = log.GetInstance()

func ExecuteNmap(c *gin.Context) {
	var request command.ApiExecuteNmap

	err := c.BindJSON(&request)
	if err == nil {
		logger.Debug("parsed request")

		target := request.Args.Target

		logger.WithField("target", target).Info("parsed request")

		if target != "" {
			nmap.Run(target)

			c.JSON(http.StatusOK, gin.H{
				"target":  target,
				"results": []string{"something here"},
			})
		}
	} else {
		logger.WithError(err).Warn("failed to parse request")
		c.JSON(http.StatusBadRequest, gin.H{"failed": true})
	}
}
