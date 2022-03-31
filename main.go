package main

import (
	"github.com/whitesource/spring4shell-detect/cmd"
	"log"
)

func main() {
	if err := cmd.NewCmdRoot().Execute(); err != nil {
		log.Fatalln(err)
	}
}
