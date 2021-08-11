package nmap

import (
	"os"
	"os/exec"
)

func runNmap(target string) {
	cmd := exec.Command("nmap", target)

	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		logger.WithError(err).Warn("failed to run nmap")
	}
}
