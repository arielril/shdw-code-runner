package main

import (
	"github.com/arielril/attack-graph-flow-code-runner/internal/router"
)

func main() {
	rt := router.CreateRouter()
	rt.Run(":5000")
}
