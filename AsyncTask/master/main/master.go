package main

import (
	"flag"
	"fmt"
	"github.com/chris2211/AsyncTask/master"
	"runtime"
	"time"
)

var (
	confFile string
)

func initArgs() {
	flag.StringVar(&confFile, "config", "./master.json", "指定master.json")
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


	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}


	if err = master.InitWorkerMgr(); err != nil {
		goto ERR
	}


	if err = master.InitLogMgr(); err != nil {
		goto ERR
	}


	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}


	if err = master.InitApiServer(); err != nil {
		goto ERR
	}


	for {
		time.Sleep(1 * time.Second)
	}

	return

ERR:
	fmt.Println(err)
}
