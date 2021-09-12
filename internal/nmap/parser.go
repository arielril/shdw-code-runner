package nmap

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/arielril/attack-graph-flow-code-runner/internal/model"
)

func NewNmapPortInfo(number int, state, protocol, owner, service, rpc_info, version string) *model.NmapPortInfo {
	return &model.NmapPortInfo{
		Number:   number,
		State:    state,
		Protocol: protocol,
		Owner:    owner,
		Service:  service,
		RpcInfo:  rpc_info,
		Version:  version,
	}
}

func ParseScanResult(resultFile string) (portResult []*model.NmapPortInfo, err error) {
	// only grepable nmap file
	fileData, err := ioutil.ReadFile(resultFile + ".gnmap")

	if err != nil {
		errMessage := "failed to parse the nmap result file"
		logger.WithError(err).Warn(errMessage)
		err = errors.New(errMessage)
		return
	}

	lines := strings.Split(string(fileData), "\n")
	logger.WithField("line_qty", len(lines)).Debug("splitted lines from result file")

	portResult = make([]*model.NmapPortInfo, 0)
	for idx, line := range lines {
		// skip lines that doesn't inform ports
		if strings.HasPrefix(line, "Host:") && strings.Contains(line, "Ports") && !strings.HasPrefix(line, "#") {
			logger.WithField("port_line", line).Debug("parsing ports from nmap")

			portList := strings.ReplaceAll(strings.Split(line, "\t")[1], "Ports: ", "")

			ports := strings.Split(portList, ",")

			for _, port := range ports {
				logger.WithField("port_data", port).Debug("parsing port result")

				portInfo := strings.Split(strings.TrimSpace(port), "/")

				portNumber, _ := strconv.Atoi(portInfo[0])
				state := portInfo[1]
				protocol := portInfo[2]
				owner := portInfo[3]
				service := portInfo[4]
				rpc_info := portInfo[5]
				version := portInfo[6]

				portResult = append(
					portResult,
					NewNmapPortInfo(portNumber, state, protocol, owner, service, rpc_info, version),
				)
			}
		} else {
			logger.WithField("port_line", line).WithField("line_index", idx).Debug("skipped line")
			continue
		}

	}

	return
}
