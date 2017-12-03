package engine

import (
	"common/log"
	"fmt"
	"time"
)

var (
	_This_Module = "Engine"
	_This_Logger *log.Logger
)

func init() {
	_This_Logger = log.New(_This_Module, "")
}

/*
功能：创建一个Engine
参数：
- 无
返回值：
- Engine: 一个Engine实例对象
*/
func New() *Engine {
	eng := &Engine{}
	eng.handlers = make(map[string]Handler)
	eng.catchall = catchall
	return eng
}

/*
功能：向Engine里注册Handler
参数：
- name: 任务(Job)名称
- handler: 负责处理任务的函数
*/
func (eng *Engine) Register(name string, handler Handler) error {
	if name == "" {
		return fmt.Errorf("Name for job can't be empty.")
	}

	if handler == nil {
		return fmt.Errorf("Handler for job can't be nil.")
	}

	if _, exists := eng.handlers[name]; exists {
		//已经存在名称相同，handler不同的Job
		return fmt.Errorf("Job %s already registered.", name)
	}

	//注册信息的Job
	eng.handlers[name] = handler
	return nil
}

func catchall(job *Job) Status {
	if job.Name == "" {
		_This_Logger.Error("The name of job is empty.")
		return StatusErr
	}
	_This_Logger.Error("No handler for job.")
	return StatusNotFound
}

/*
功能：创建一个Job
参数：
- name: Job的名称
- args: Job执行需要的参数
返回值：
- Job: 一个Job实例对象
*/
func (eng *Engine) Job(name string, args ...string) *Job {
	job := &Job{
		Eng:  eng,
		Name: name,
		Args: args,
	}

	if handler, exists := eng.handlers[name]; exists {
		job.handler = handler
	} else if eng.catchall != nil {
		job.handler = eng.catchall
	}

	return job
}

//判断Engine是否关闭了
func (eng *Engine) IsShutdown() bool {
	return eng.shutdown
}

func (eng *Engine) Shutdown() {
	//加写锁，设置engine状态为关闭
	eng.l.Lock()
	if eng.shutdown {
		eng.l.Unlock()
		return
	}
	eng.shutdown = true
	eng.l.Unlock()

	//等待所有任务结束，设置任务超时时间为10秒
	tasksDone := make(chan struct{})
	go func() {
		eng.tasks.Wait()
		close(tasksDone)
	}()

	select {
	case <-time.After(time.Second * 10):
	case <-tasksDone:
	}
}
