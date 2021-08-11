package nmap

import (
	"os"
	"os/exec"

	"github.com/arielril/attack-graph-flow-code-runner/log"
)

var logger = log.GetInstance()

func Run(target string) {
	cmd := exec.Command("nmap", target)

	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		logger.WithError(err).Warn("failed to run nmap")
	}
}
