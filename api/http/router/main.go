package router

import (
	"os"

	"github.com/arielril/attack-graph-flow-code-runner/api/http/nmapapi"
	"github.com/arielril/attack-graph-flow-code-runner/log"
	"github.com/gin-gonic/gin"
)

var logger = log.GetInstance()

func CreateRouter() *gin.Engine {
	rt := gin.Default()

	rt.POST("/v1/nmap", nmapapi.ExecuteNmap)

	return rt
}

func SetupOutputFolder() {
	tmpDir, err := os.MkdirTemp("", "")
	if err != nil {
		logger.WithError(err).Error("failed to initialize temp dir")
		return
	}

	logger.WithField("temp_dir", tmpDir).Info("temp dir created")

	os.Setenv("xxxx_TMP_DIR", tmpDir)

	// create NMAP folder
	nmapDir, _ := os.MkdirTemp(tmpDir, "nmap")
	os.Setenv("xxxx_TMP_DIR_NMAP", nmapDir)
	// create FFuF folder
	os.MkdirTemp(tmpDir, "ffuf")
}
