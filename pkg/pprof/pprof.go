package pprof

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
)

const profFile = "out.prof"

var profile = flag.Bool("profile", false, fmt.Sprintf("Dump pprof to %s", profFile))

func MaybeProfile() {
	flag.Parse()
	if *profile {
		f, err := os.Create("out.prof")
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
	}
}

func MaybeStopProfile() {
	if *profile {
		pprof.StopCPUProfile()
	}
}
