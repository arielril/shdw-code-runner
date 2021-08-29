package main

import "github.com/arielril/attack-graph-flow-code-runner/api/http/router"

func main() {
	rt := router.CreateRouter()

	router.SetupOutputFolder()

	rt.Run(":5000")
}
