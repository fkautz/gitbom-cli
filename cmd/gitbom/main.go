package main

import (
	gitbom "github.com/fkautz/gitbom-cli/pkg/cmd"
	"log"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	gitbom.Cmd.Run()
}
