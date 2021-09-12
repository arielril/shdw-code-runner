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
	var request model.ApiExecuteNmap

	err := c.BindJSON(&request)
	if err == nil {
		logger.Debug("parsed request")

		nmapOpts := nmap.ApiExecuteToInternalStruct(request)

		logger.WithField("target", nmapOpts.Target).Info("parsed request")

		if nmapOpts.Target != "" {
			resultFile, err := nmap.Run(nmapOpts)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   err,
					"message": "failed to run nmap",
				})
				return
			}

			portResult, err := nmap.ParseScanResult(resultFile)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   err,
					"message": "failed to parse result from nmap",
				})
				return
			}

			response := make([]model.ApiExecuteNmap200PortsResponse, 0)

			for _, port := range portResult {
				response = append(
					response,
					model.ApiExecuteNmap200PortsResponse{
						Number:   port.Number,
						State:    port.State,
						Protocol: port.Protocol,
						Owner:    port.Owner,
						Service:  port.Service,
						RpcInfo:  port.RpcInfo,
						Version:  port.Version,
					},
				)
			}

			resp := &model.ApiExecuteNmap200Response{
				Target: nmapOpts.Target,
				Ports:  response,
			}

			c.JSON(http.StatusOK, resp)
		}
	} else {
		logger.WithError(err).Warn("failed to parse request")
		c.JSON(http.StatusBadRequest, gin.H{"failed": true})
	}
}
