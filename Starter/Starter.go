package Starter

import (
	"../DaemonConfig"
	"../Logger"
	"log"
	"os/exec"
	"strings"
)

func Operate(daemon DaemonConfig.Daemon) {
	commandParts := strings.Fields(daemon.Cmd)
	if len(commandParts) < 2 {
		panic("Can't parse exec command")
	}

	Logger.WriteToLog(daemon.LogFile, "start: "+daemon.Cmd)

	cmd := exec.Command(commandParts[0], commandParts[1:]...)
	output, err2 := cmd.CombinedOutput()
	if err2 != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err2)
	}

	Logger.WriteToLog(daemon.LogFile, "result: "+string(output))
}
