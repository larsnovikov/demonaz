package main

import (
	"./DaemonConfig"
	"./Starter"
	"fmt"
)

func main() {
	config := DaemonConfig.GetConfig()
	for i := 0; i < len(config.Daemons); i++ {
		if config.Daemons[i].Enable == true {
			for j := 1; j <= config.Daemons[i].WorkerCount; j++ {
				go startRoutine(config.Daemons[i])
			}
		}
	}
	var input string
	fmt.Scanln(&input)
}

func startRoutine(daemon DaemonConfig.Daemon) {
	for {
		Starter.Operate(daemon)
	}
}
