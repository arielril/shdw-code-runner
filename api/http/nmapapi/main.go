package nmapapi

import (
	"net/http"

	"github.com/arielril/attack-graph-flow-code-runner/internal/model"
	"github.com/arielril/attack-graph-flow-code-runner/internal/nmap"
	"github.com/arielril/attack-graph-flow-code-runner/log"
	"github.com/gin-gonic/gin"
)

var logger = log.GetInstance()

func ExecuteNmap(c *gin.Context) {
	var request *model.ApiExecuteNmap

	err := c.BindJSON(&request)
	if err == nil {
		logger.Debug("parsed request")

		nmapOpts := nmap.ApiExecuteToInternalStruct(*request)

		logger.WithField("target", nmapOpts.Target).Info("parsed request")

		if nmapOpts.Target != "" {
			nmap.Run(nmapOpts)

			resp := &model.ApiExecuteNmap200Response{
				Target: nmapOpts.Target,
				Ports: []*model.ApiExecuteNmap200PortsResponse{
					{
						Number:   80,
						State:    "open",
						Protocol: "TCP",
						Service:  "http",
						Version:  "Apache httpd 2.4.41 ((Ubuntu))",
					},
				},
			}

			c.JSON(http.StatusOK, resp)
		}
	} else {
		logger.WithError(err).Warn("failed to parse request")
		c.JSON(http.StatusBadRequest, gin.H{"failed": true})
	}
}
