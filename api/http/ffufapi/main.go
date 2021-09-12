package ffufapi

import (
	"net/http"

	"github.com/arielril/attack-graph-flow-code-runner/internal/ffuf"
	"github.com/arielril/attack-graph-flow-code-runner/internal/model"
	"github.com/arielril/attack-graph-flow-code-runner/log"
	"github.com/gin-gonic/gin"
)

var logger = log.GetInstance()

func ExecuteFfuf(c *gin.Context) {
	var request model.ApiExecuteFfuf

	err := c.BindJSON(&request)
	if err != nil {
		logger.WithError(err).Warn("failed to parse request")
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse request"})
		return
	}

	logger.
		WithField("api", "ffuf").
		WithField("operation", "execute").
		Debug("parsed request")

	ffufOpts := ffuf.ApiExecuteToInternalStruct(request)

	if ffufOpts.Target == "" {
		logger.Warn("received request without a target")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid target, empty",
		})
		return
	}

	resultFile, err := ffuf.Run(ffufOpts)
	if err != nil {
		errMessage := "failed to execute ffuf and get a result"
		logger.WithError(err).Warn(errMessage)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err,
			"message": errMessage,
		})
		return
	}

	results, err := ffuf.ParseResultFile(resultFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	response := ffuf.FfufOutputResultToApiExecute200Response(ffufOpts.Target, results)
	c.JSON(http.StatusOK, response)
}
