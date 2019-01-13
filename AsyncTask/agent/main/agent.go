package main

import (
	"flag"
	"fmt"
	"github.com/chris2211/AsyncTask/agent"
	"agent"
	"runtime"
	"time"
)

var (
	confFile string
)

func initArgs() {
	// agent -config ./worker.json
	// agent -h
	flag.StringVar(&confFile, "config", "./agent.json", "agent.json")
	flag.Parse()
}

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)

	initArgs()

	initEnv()

	if err = agent.InitConfig(confFile); err != nil {
		goto ERR
	}

	if err = agent.InitRegister(); err != nil {
		goto ERR
	}

	if err = agent.InitLogSink(); err != nil {
		goto ERR
	}

	if err = agent.InitExecutor(); err != nil {
		goto ERR
	}

	if err = agent.InitScheduler(); err != nil {
		goto ERR
	}

	if err = agent.InitJobMgr(); err != nil {
		goto ERR
	}

	for {
		time.Sleep(1 * time.Second)
	}

	return

ERR:
	fmt.Println(err)
}
