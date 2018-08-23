package Starter

import (
	"../DaemonConfig"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Operate(daemon DaemonConfig.Daemon) {
	commandParts := strings.Fields(daemon.Cmd)
	if len(commandParts) < 2 {
		fmt.Println("shit happens")
	}

	cmd := exec.Command(commandParts[0], commandParts[1:]...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	// fmt.Printf(string(out))
}
