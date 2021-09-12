package nmap

import (
	"os"
	"os/exec"

	"github.com/arielril/attack-graph-flow-code-runner/internal/model"
	"github.com/arielril/attack-graph-flow-code-runner/internal/util/rnd"
	"github.com/arielril/attack-graph-flow-code-runner/log"
)

var logger = log.GetInstance()

func Run(opts *model.NmapOptions) (outFile string, err error) {
	var outputFolder = os.Getenv("xxxx_TMP_DIR_NMAP")

	outFile = outputFolder + "/" + rnd.StringWithLength(5) + opts.Target
	cmd := exec.Command("nmap", opts.Target, "-oA", outFile)

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

	if opts.Verbose {
		cmd.Args = append(cmd.Args, "-v")
	}

	err = cmd.Run()
	if err != nil {
		logger.WithError(err).Warn("failed to run nmap")
	}

	return
}
