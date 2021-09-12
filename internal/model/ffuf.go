package model

type ApiExecuteFfuf struct {
	Target       string `json:"target"`
	Recursion    bool   `json:"recursion" default:"false"`
	Redirect     bool   `json:"redirect" default:"false"`
	IgnoreStatus []int  `json:"ignore_status"`
}

type ApiExecuteFfuf200Response struct {
	Target string                          `json:"target"`
	Data   []ApiExecuteFfuf200PathResponse `json:"data"`
}

type ApiExecuteFfuf200PathResponse struct {
	StatusCode int                                  `json:"status_code"`
	Path       string                               `json:"path"`
	URL        string                               `json:"url"`
	Content    ApiExecuteFfuf200PathContentResponse `json:"content"`
}

type ApiExecuteFfuf200PathContentResponse struct {
	Length int `json:"length,omitempty"`
	Words  int `json:"words,omitempty"`
	Lines  int `json:"lines,omitempty"`
}

type FfufOptions struct {
	Target       string
	Recursion    bool
	Redirect     bool
	FilterStatus string
}

type FfufOutputResult struct {
	Result []FfufOutputResultItem `json:"results"`
}

type FfufOutputResultItem struct {
	Input    map[string]string `json:"input"`
	Position int               `json:"position"`
	Status   int               `json:"status"`
	Length   int               `json:"length"`
	Words    int               `json:"word"`
	Lines    int               `json:"lines"`
	URL      string            `json:"url"`
}
