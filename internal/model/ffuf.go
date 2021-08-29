package model

type ApiExecuteFfuf struct {
	Target       string `json:"target"`
	Recursion    bool   `json:"recursion" default:"false"`
	Redirect     bool   `json:"redirect" default:"false"`
	IgnoreStatus []int  `json:"ignore_status"`
}

type ApiExecuteFfuf200Response struct {
	Data []ApiExecuteFfuf200PathResponse `json:"data"`
}

type ApiExecuteFfuf200PathResponse struct {
	StatusCode int                                  `json:"status_code"`
	Path       string                               `json:"path"`
	Content    ApiExecuteFfuf200PathContentResponse `json:"content"`
}

type ApiExecuteFfuf200PathContentResponse struct {
	Length int `json:"length,omitempty"`
	Words  int `json:"words,omitempty"`
	Lines  int `json:"lines,omitempty"`
}
