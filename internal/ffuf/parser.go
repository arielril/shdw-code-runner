package ffuf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/arielril/attack-graph-flow-code-runner/internal/model"
)

func ParseResultFile(outputFile string) (result model.FfufOutputResult, err error) {
	fileData, err := ioutil.ReadFile(outputFile)
	if err != nil {
		errMessage := "failed to read ffuf output file"
		logger.WithError(err).Warn(errMessage)
		return
	}

	err = json.Unmarshal(fileData, &result)
	if err != nil {
		errMessage := "failed to parse ffuf output file"
		logger.WithError(err).Warn(errMessage)
		return
	}

	return
}
