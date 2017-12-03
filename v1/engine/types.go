package engine

import (
	"sync"
	"time"
)

type Status int8
type Handler func(*Job) Status

const (
	StatusOk       Status = 0   //任务成功
	StatusErr      Status = 1   //任务失败
	StatusNotFound Status = 127 //任务没执行
)

type Engine struct {
	handlers map[string]Handler
	catchall Handler
	tasks    sync.WaitGroup
}

type Job struct {
	Eng     *Engine
	Name    string
	Args    []string
	id      string //随机字符串
	handler Handler
	status  Status
	end     time.Time
}

const (
	_ERR_HANDLER_IS_NULL string = "The handler of this job is null."
)
