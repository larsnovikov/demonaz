package main

import (
	"./DaemonConfig"
	"./Starter"
	"fmt"
	"github.com/marcsauter/single"
	"log"
	"os"
	"sync"
	"time"
)

var checkConfigChanges bool

func main() {
	s := single.New("demonaz")
	if err := s.CheckLock(); err != nil && err == single.ErrAlreadyRunning {
		log.Fatal("another instance of the app is already running, exiting")
	} else if err != nil {
		log.Fatalf("failed to acquire exclusive app lock: %v", err)
	}
	defer s.TryUnlock()

	args := os.Args[1:]
	if len(args) == 0 {
		panic("Need config file as arg")
	}
	config := DaemonConfig.GetConfig(args[0])
	checkConfigChanges = config.CheckConfigChanges

	var wg sync.WaitGroup

	for i := 0; i < len(config.Daemons); i++ {
		if config.Daemons[i].Enable == true {
			for j := 1; j <= config.Daemons[i].WorkerCount; j++ {
				wg.Add(1)
				time.Sleep(time.Duration(config.Daemons[i].StartDelay) * time.Second)
				go startRoutine(&wg, config.Daemons[i])
			}
		}
	}

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")
}

func startRoutine(wg *sync.WaitGroup, daemon DaemonConfig.Daemon) {
	defer wg.Done()
	for {
		if checkConfigChanges && DaemonConfig.IsConfigChanged() {
			fmt.Println("Config changed! Stop goroutines")
			break
		}
		Starter.Operate(daemon)
	}
}
