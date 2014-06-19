package main

import (
	"fmt"
	"github.com/dotcloud/docker/pkg/system"
	"os/exec"
	"syscall"
	"os"
)

func main() {
	_, slave, err := system.CreateMasterAndConsole()
	if err != nil {
		fmt.Println("CreateMasterAndConsole failed: ", err)
		return
	}

	fmt.Println("os.Args[0] = ", os.Args[0])

	command := exec.Command(os.Args[1], slave)
	system.SetCloneFlags(command, uintptr(syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNS |
				syscall.CLONE_NEWPID | syscall.CLONE_NEWNET))

	if output, err := command.Output(); err == nil {
		fmt.Println("command.Output: ", string(output[:]))
	} else {
		fmt.Println("command.Output error = ", err)
		return
	}

	/*_, err = system.OpenTerminal(slave, syscall.O_RDWR)
	if err != nil {
		fmt.Println("OpenTerminal failed: ", err)
		return
	}*/
	fmt.Println("testpts succeeded")
}
