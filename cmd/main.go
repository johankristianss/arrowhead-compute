package main

import (
	"github.com/johankristianss/arrowhead-compute/internal/cli"
	"github.com/johankristianss/arrowhead-compute/pkg/build"
)

var (
	BuildVersion string = ""
	BuildTime    string = ""
)

func main() {
	build.BuildVersion = BuildVersion
	build.BuildTime = BuildTime
	cli.Execute()
}
