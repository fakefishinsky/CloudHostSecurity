package log

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	_LOG_ALERT string = "[ALERT]\t"
	_LOG_ERROR string = "[ERROR]\t"
	_LOG_WARN  string = "[WARNING]\t"
	_LOG_INFO  string = "[INFO]\t"
)

type Logger struct {
	Name    string
	LogFile string
	logFd   *os.File
	mtx     sync.Mutex
}

type LoggerResource struct {
	loggers map[string]*Logger
	mtx     sync.Mutex
}

var (
	_Logger_Resource LoggerResource
)

func init() {
	_Logger_Resource = LoggerResource{}
	_Logger_Resource.loggers = make(map[string]*Logger)
}

//从资源池中获取名为name的Logger
func Get(name string) *Logger {
	_Logger_Resource.mtx.Lock()
	defer _Logger_Resource.mtx.Unlock()

	return _Logger_Resource.loggers[name]
}

//从资源池中删除名为name的Logger
func Remove(name string) {
	_Logger_Resource.mtx.Lock()
	defer _Logger_Resource.mtx.Unlock()

	if logger, exists := _Logger_Resource.loggers[name]; exists {
		if logger.logFd != nil {
			logger.logFd.Close()
		}
		delete(_Logger_Resource.loggers, name)
	}
}

//Logger初始化
func New(name string, logFile string) *Logger {
	_Logger_Resource.mtx.Lock()
	defer _Logger_Resource.mtx.Unlock()

	if logger, exists := _Logger_Resource.loggers[name]; exists {
		return logger
	}

	logger := &Logger{Name: name, LogFile: logFile}
	if logFile != "" {
		logger.logFd, _ = os.OpenFile(logFile, os.O_APPEND|os.O_CREATE, 0600)
	}

	_Logger_Resource.loggers[name] = logger
	return logger
}

//记录日志
func (l *Logger) record(msg string, tag string) {
	l.mtx.Lock() //加锁，避免日志混乱

	tm := time.Now()
	logMsg := fmt.Sprintf("%s: %d-%02d-%02d %02d:%02d:%02d %s%s\n", l.Name, tm.Year(), int(tm.Month()), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), tag, msg)
	if l.logFd != nil {
		//写到日志文件里
		l.logFd.WriteString(logMsg)
	} else {
		//打印到控制台
		fmt.Print(logMsg)
	}

	l.mtx.Unlock()
}

func (l *Logger) Alert(msg string) {
	l.record(msg, _LOG_ALERT)
}

func (l *Logger) Error(msg string) {
	l.record(msg, _LOG_ERROR)
}

func (l *Logger) Warn(msg string) {
	l.record(msg, _LOG_WARN)
}

func (l *Logger) Info(msg string) {
	l.record(msg, _LOG_INFO)
}
