package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "server":
		runServer()
	case "test":
		runTest()
	default:
		help()
	}
}

func help() {
	fmt.Println("Usage:community <command> [options]")
	fmt.Println("commands:")
	fmt.Println("  server [-min]")
}
