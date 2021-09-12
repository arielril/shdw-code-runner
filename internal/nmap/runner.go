package nmap

import (
	"os"
	"os/exec"

	"github.com/arielril/attack-graph-flow-code-runner/internal/model"
	"github.com/arielril/attack-graph-flow-code-runner/internal/util/rnd"
	"github.com/arielril/attack-graph-flow-code-runner/log"
	"github.com/sirupsen/logrus"
)

var logger = log.GetInstance()

func Run(opts *model.NmapOptions) (outFile string, err error) {
	var outputFolder = os.Getenv("xxxx_TMP_DIR_NMAP")

	outFile = outputFolder + "/" + rnd.StringWithLength(5) + opts.Target
	cmd := exec.Command("nmap", opts.Target, "-oA", outFile)

	logOpts := logrus.Fields{
		"target": opts.Target,
	}

	if opts.DefaultScripts {
		cmd.Args = append(cmd.Args, "-sC")
		logOpts["default_scripts"] = true
	}

	if opts.ServiceVersion {
		cmd.Args = append(cmd.Args, "-sV")
		logOpts["service_version"] = true
	}

	if opts.HostsOnline {
		cmd.Args = append(cmd.Args, "-Pn")
		logOpts["hosts_online"] = true
	}

	if opts.PortRange != "" {
		cmd.Args = append(cmd.Args, "-p", opts.PortRange)
		logOpts["port_range"] = opts.PortRange
	}

	if opts.SynScan {
		cmd.Args = append(cmd.Args, "-sS")
		logOpts["syn_scan"] = true
	}

	if opts.UdpScan {
		cmd.Args = append(cmd.Args, "-sU")
		logOpts["udp_scan"] = true
	}

	if opts.Verbose {
		cmd.Args = append(cmd.Args, "-v")
		logOpts["verbose"] = true
	}

	logger.WithFields(logOpts).Info("running nmap")
	err = cmd.Run()
	if err != nil {
		logger.WithError(err).Warn("failed to run nmap")
	} else {
		logger.WithField("file", outFile).Debug("nmap result wrote to file")
	}

	return
}
