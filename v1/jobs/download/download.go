package download

import (
	"engine"
	"io/ioutil"
	"net"
	"net/http"
)

const (
	JOB_NAME string = "download"
)

/*
功能:download任务入口
参数:
- job: Job结构实例指针
       job.Name = "download"
	   job.Args = ["uri(must)", "file-path(must)", "unix(optional)"]
*/
func JobEntry(job *engine.Job) engine.Status {
	if job.Name != JOB_NAME {
		return engine.StatusNotFound
	}

	argsLen := len(job.Args)
	if argsLen < 2 {
		return engine.StatusErr
	}

	uri := job.Args[0]
	file := job.Args[1]
	flUnix := false
	if argsLen >= 3 && job.Args[2] == "unix" {
		flUnix = true
	}

	if err := doDownload(uri, file, flUnix); err != nil {
		return engine.StatusErr
	}

	return engine.StatusOk
}

func doDownload(uri, file string, flUnix bool) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if flUnix {
		//将数据写到unix socket中, 适合少量数据

		//解析地址
		unixAddr, err := net.ResolveUnixAddr("unix", file)
		if err != nil {
			return err
		}

		//获取连接
		unixConn, err := net.DialUnix("unix", nil, unixAddr)
		if err != nil {
			return err
		}

		//写数据
		err = unixConn.Write(data)
		if err != nil {
			unixConn.Close()
			return err
		}

		return unixConn.Close()
	} else {
		//将数据写到普通文件中, 适合大量数据
		return ioutil.WriteFile(file, data, 0600)
	}
}
