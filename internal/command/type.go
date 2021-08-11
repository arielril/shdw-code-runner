package command

type ApiExecuteNmap struct {
	Args struct {
		Target string `json:"target"`
	} `json:"args"`
}
