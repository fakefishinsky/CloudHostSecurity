package engine

import (
	"fmt"
	"time"
)

//运行任务
func (job *Job) Run() error {
	//加写锁，添加任务
	job.Eng.l.Lock()
	if job.Eng.IsShutdown() {
		//engine已经关闭，任务结束
		job.Eng.l.Unlock()
		return fmt.Errorf("Engine is shutdown.")
	} else {
		//添加任务
		job.Eng.tasks.Add(1)
		job.Eng.l.Unlock()
		defer job.Eng.tasks.Done()
	}

	if job.handler != nil {
		job.status = job.handler(job)
		job.end = time.Now()
		return nil
	}

	job.status = StatusNotFound
	return fmt.Errorf("No handler found for job(%s).", job.Name)
}

//获取任务状态
func (job *Job) Status() string {
	if job.end.IsZero() {
		//任务未执行或者尚未运行完
		return ""
	}

	stat := "ERR"
	if job.status == StatusOk {
		stat = "OK"
	}

	return fmt.Sprintf("%s (%d)", stat, job.status)
}
