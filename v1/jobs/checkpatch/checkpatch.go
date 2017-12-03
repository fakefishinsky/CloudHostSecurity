package checkpatch

import (
	"bytes"
	"common"
	"engine"
	"fmt"
	"os/exec"
	"strings"
)

const (
	JOB_NAME     string = "checkpatch"
	_THIS_MODULE string = "checkpatch"
)

var (
	_Sys_Pkgs *SystemPackage
)

func init() {
	_Sys_Pkgs = new(SystemPackage)
	_Sys_Pkgs.packages = make(map[string]Version)
}

func JobEntry(job *engine.Job) engine.Status {
	if job.Name != JOB_NAME {
		return engine.StatusNotFound
	}

	fmt.Println("Checking Security Patch")
	_Sys_Pkgs.packages["openssh"] = Version{"6.2P1", "0.40.1"}
	fmt.Println(len(_Sys_Pkgs.packages))
	return engine.StatusOk
}

//获取系统软件包信息
func (sysPkgs *SystemPackage) query() error {
	//执行rpm -qa命令获取系统已安装的软件包
	resBuf := bytes.Buffer{}
	cmd := exec.Command("rpm", "-qa")
	cmd.Stdout = &resBuf
	err := cmd.Run()
	if err != nil {
		return err
	}
	if resBuf.Len() <= 0 {
		return common.Error{_ERR_CMDOUT_IS_NULL}
	}

	//加写锁
	sysPkgs.rwmtx.Lock()

	//解析命令输出的结果
	pkgs := strings.Split(resBuf.String(), "\n")
	for _, pkg := range pkgs {
		name, version, release := getPackageInfo(pkg)
		if name != "" && version != "" {
			sysPkgs.packages[name] = Version{version, release}
		}
	}

	sysPkgs.rwmtx.Unlock()
	return nil
}

//获取一个软件包的名称和版本信息(name-version-release)
func getPackageInfo(pkg string) (name, version, release string) {
	name, version, release = "", "", ""

	pkgInfos := strings.Split(pkg, "-")

	//至少要包含2个信息(name-version)
	infoLen := len(pkgInfos)
	if infoLen < 2 {
		return name, version, release
	}

	//删除软件包后面的格式信息
	last := infoLen - 1
	for _, s := range []string{".rpm", ".x86_64", ".noarch", ".src", ".nosrc"} {
		pkgInfos[last] = strings.TrimRight(pkgInfos[last], s)
	}

	if infoLen == 2 {
		name = pkgInfos[0]
		version = pkgInfos[1]
	} else {
		name = strings.Join(pkgInfos[:infoLen-2], "-")
		version = pkgInfos[infoLen-2]
		release = pkgInfos[infoLen-1]
	}
	return name, version, release
}

/*
功能:比较版本高低
参数：
- v1, v2:两个要比较的版本号
返回值:
- =0: v1 = v2
- >0: v1 > v2
- <0: v1 < v2
*/
func compareVersion(v1, v2 string) int {
	return 0
}

func splitVersion(ver string) []string {
	return []string{}
}

//单组补丁检查
func (ph *Patch) check() {

}

//补丁集检查
func (ps *Patchset) check() {

}
