package util

import (
	"bytes"
	"os/exec"
)

//阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
// @s 执行的命令
func CMD(s string) (string, string) {
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout
	//再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	var eb bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &eb

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	cmd.Run()
	return out.String(), eb.String()
}
