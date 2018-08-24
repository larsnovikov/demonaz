package main

import (
	"./DaemonConfig"
	"./Starter"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		panic("Need config file as arg")
	}
	config := DaemonConfig.GetConfig(args[0])
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
