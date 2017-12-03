package engine

import (
	"common"
	"common/log"
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
	eng.handlers = map[string]Handler{}
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

//运行任务
func (job *Job) Run() error {
	defer job.Eng.tasks.Done()

	if job.handler != nil {
		job.status = job.handler(job)
		job.end = time.Now()
		return nil
	}

	return common.Error{_ERR_HANDLER_IS_NULL}
}

func (eng *Engine) AddTask(n int) {
	eng.tasks.Add(n)
}

func (eng *Engine) Wait() {
	eng.tasks.Wait()
}
