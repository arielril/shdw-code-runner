package nmap

import (
	"os"
	"os/exec"

	"github.com/arielril/attack-graph-flow-code-runner/internal/model"
	"github.com/arielril/attack-graph-flow-code-runner/log"
)

var logger = log.GetInstance()

func Run(opts *model.NmapOptions) {
	var outputFolder = os.Getenv("xxxx_TMP_DIR_NMAP")

	outFileName := outputFolder + "/" + opts.Target
	cmd := exec.Command("nmap", opts.Target, "-oA", outFileName)

	if opts.DefaultScripts {
		cmd.Args = append(cmd.Args, "-sC")
	}

	if opts.ServiceVersion {
		cmd.Args = append(cmd.Args, "-sV")
	}

	if opts.HostsOnline {
		cmd.Args = append(cmd.Args, "-Pn")
	}

	if opts.PortRange != "" {
		cmd.Args = append(cmd.Args, "-p", opts.PortRange)
	}

	if opts.SynScan {
		cmd.Args = append(cmd.Args, "-sS")
	}

	if opts.UdpScan {
		cmd.Args = append(cmd.Args, "-sU")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		logger.WithError(err).Warn("failed to run nmap")
	}
}
