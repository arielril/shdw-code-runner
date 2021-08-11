package main

import "github.com/arielril/attack-graph-flow-code-runner/api/http/router"

func main() {
	rt := router.CreateRouter()
	rt.Run(":5000")
}
