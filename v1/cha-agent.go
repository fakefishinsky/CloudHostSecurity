package main

import (
	"engine"
	"fmt"
	"jobs/checkconfig"
	"jobs/checkpatch"
	"jobs/download"
	"runtime"
)

func test() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	eng := engine.New()
	eng.Register(checkpatch.JOB_NAME, checkpatch.JobEntry)
	eng.Register(checkconfig.JOB_NAME, checkconfig.JobEntry)
	eng.Register(download.JOB_NAME, download.JobEntry)

	eng.AddTask(3)
	go eng.Job("checkconfig", "hello").Run()
	go eng.Job("checkpatch").Run()
	go eng.Job("download", "http://www.baidu.com/", "baidu.html").Run()
	eng.Wait()
}

func main() {
	fmt.Println("Test begin.")
	test()
	fmt.Println("Test finished.")
}
