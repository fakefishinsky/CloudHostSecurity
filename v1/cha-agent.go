package main

import (
	"engine"
	"fmt"
	"jobs/checkconfig"
	"jobs/checkpatch"
	"jobs/download"
	"runtime"
	"time"
)

func test() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	eng := engine.New()
	eng.Register(checkpatch.JOB_NAME, checkpatch.JobEntry)
	eng.Register(checkconfig.JOB_NAME, checkconfig.JobEntry)
	eng.Register(download.JOB_NAME, download.JobEntry)

	go eng.Job("checkconfig", "hello").Run()
	go eng.Job("checkpatch").Run()
	go eng.Job("download", "http://www.baidu.com/", "baidu.html").Run()
	
	time.Sleep(time.Second * 3)
	eng.Shutdown()
}

func main() {
	fmt.Println("Test begin.")
	test()
	fmt.Println("Test finished.")
}
