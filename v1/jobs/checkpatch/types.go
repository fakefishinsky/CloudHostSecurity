package checkpatch

import (
	"sync"
)

type CheckResult uint

const ( //单组补丁检查结果
	_CR_SKIP  CheckResult = 0x1    //不涉及，跳过
	_CR_PASS  CheckResult = 0x10   //检查通过
	_CR_FAIL  CheckResult = 0x100  //检查不通过
	_CR_ERROR CheckResult = 0x1000 //检查出错
)

const ( //error信息
	_ERR_CMDOUT_IS_NULL string = "The result of command(rpm -qa) is empty."
)

//补丁信息结构
type Patch struct {
	Id          string            //uuid
	Date        string            //补丁发布日期
	Abstract    string            //补丁摘要(标题)
	Packages    []string          //补丁包名
	CVEs        map[string]string //解决的CVE漏洞以及CVSS评分
	Description string            //问题官方描述
	OfficialUrl string            //补丁官方链接
	insPackages []string          //系统已安装的包
	result      CheckResult       //检查结果
}

//补丁集结构
type Patchset struct {
	Version string  //补丁集版本
	Total   int     //补丁总量
	Patches []Patch //补丁信息集合
}

//软件包版本
type Version struct {
	version string
	release string
}

type SystemPackage struct {
	packages map[string]Version
	rwmtx    sync.RWMutex
}
