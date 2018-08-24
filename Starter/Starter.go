package Starter

import (
	"../DaemonConfig"
	"log"
	"os/exec"
	"strings"
)

func Operate(daemon DaemonConfig.Daemon) {
	commandParts := strings.Fields(daemon.Cmd)
	if len(commandParts) < 2 {
		panic("Can't parse exec command")
	}

	cmd := exec.Command(commandParts[0], commandParts[1:]...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
