package Starter

import (
	"../daemon_config"
	"../logger"
	"fmt"
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
	output, err := cmd.CombinedOutput()
	if err != nil {
		Logger.WriteToLog(daemon.LogFile, fmt.Sprintf("cmd.Run() failed with %s\n", err.Error()))
	}

	Logger.WriteToLog(daemon.LogFile, "result: "+string(output))
}
