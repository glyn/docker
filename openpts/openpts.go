package main

import (
	"fmt"
	"github.com/dotcloud/docker/pkg/system"
	"syscall"
	"os"
)

func main() {
	slave := os.Args[1]

	_, err := system.OpenTerminal(slave, syscall.O_RDWR)
	if err != nil {
		fmt.Println("OpenTerminal failed: ", err)
		return
	}
	fmt.Println("openpts succeeded")
}
