package main

import "dnd-agent/pkg/pipeline"

func main() {
	e := pipeline.NewEngine()
	e.Run()
}
