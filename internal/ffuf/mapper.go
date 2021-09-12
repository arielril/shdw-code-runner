package ffuf

import (
	"strconv"
	"strings"

	"github.com/arielril/attack-graph-flow-code-runner/internal/model"
)

func ApiExecuteToInternalStruct(apiEx model.ApiExecuteFfuf) *model.FfufOptions {
	r := new(model.FfufOptions)

	r.Target = apiEx.Target
	r.Recursion = apiEx.Recursion
	r.Redirect = apiEx.Redirect

	if apiEx.IgnoreStatus != nil && len(apiEx.IgnoreStatus) > 0 {
		statusList := make([]string, 0)
		for _, status := range apiEx.IgnoreStatus {
			statusList = append(statusList, strconv.Itoa(status))
		}
		r.FilterStatus = strings.Join(statusList, ",")
	}

	return r
}

func FfufOutputResultToApiExecute200Response(target string, output model.FfufOutputResult) model.ApiExecuteFfuf200Response {
	r := model.ApiExecuteFfuf200Response{
		Target: target,
	}

	if output.Result != nil && len(output.Result) > 0 {
		r.Data = make([]model.ApiExecuteFfuf200PathResponse, 0)
		for _, result := range output.Result {
			r.Data = append(
				r.Data,
				model.ApiExecuteFfuf200PathResponse{
					StatusCode: result.Status,
					Path:       result.Input["FUZZ"],
					URL:        result.URL,
					Content: model.ApiExecuteFfuf200PathContentResponse{
						Length: result.Length,
						Words:  result.Words,
						Lines:  result.Lines,
					},
				},
			)
		}
	}

	return r
}
