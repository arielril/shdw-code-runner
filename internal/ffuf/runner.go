package ffuf

import (
	"os"
	"os/exec"

	"github.com/arielril/attack-graph-flow-code-runner/internal/model"
	"github.com/arielril/attack-graph-flow-code-runner/internal/util/rnd"
	"github.com/arielril/attack-graph-flow-code-runner/log"
	"github.com/sirupsen/logrus"
)

const defaultWordlist = "/opt/seclists/Discovery/Web-Content/raft-small-words.txt"

var logger = log.GetInstance()

func Run(opts *model.FfufOptions) (outputFile string, err error) {
	outputFolder := os.Getenv("xxxx_TMP_DIR_FFUF")
	outputFile = outputFolder + "/" + rnd.StringWithLength(7)

	cmd := exec.Command(
		"ffuf",
		"-u", opts.Target,
		"-w", defaultWordlist,
		"-o", outputFile,
	)

	logOpts := logrus.Fields{
		"target": opts.Target,
	}

	if opts.FilterStatus != "" {
		cmd.Args = append(cmd.Args, "-fc", opts.FilterStatus)
		logOpts["filter_status"] = opts.FilterStatus
	}

	if opts.Recursion {
		cmd.Args = append(cmd.Args, "-recursion")
		logOpts["recursion"] = true
	}

	if opts.Redirect {
		cmd.Args = append(cmd.Args, "-r")
		logOpts["recursion"] = true
	}

	logger.WithFields(logOpts).Info("running ffuf")
	err = cmd.Run()
	if err != nil {
		logger.WithError(err).Warn("failed to execute FFuf")
	} else {
		logger.
			WithField("command", "ffuf").
			WithField("file", outputFile).
			Info("ffuf result wrote to file")
	}

	return
}
