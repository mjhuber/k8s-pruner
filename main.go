package main

import (
	"github.com/mjhuber/k8s-pruner/cmd"
)

var (
	// VERSION is set during build
	VERSION = "dev"
)

func main() {
	cmd.Execute(VERSION)
}
