package cli

import (
	"flag"

	"github.com/neurader/nqs/internal/process"
)

func Run() {
	name := flag.String("n", "", "process name")
	del := flag.String("d", "", "delete process")
	flag.Parse()

	args := flag.Args()

	if *del != "" {
		process.Delete(*del)
		return
	}

	if *name != "" && len(args) > 0 {
		process.Start(*name, args)
		return
	}

	process.List()
}
