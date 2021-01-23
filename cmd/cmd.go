package cmd

import (
	"bytes"
	"fmt"
	"github.com/PeterYangs/tools"
	"io"
	"log"
	"os/exec"
	"runtime"
)

//命令输出日志
type WriteLog struct {
	Cmd  string
	Name string
}

//命令报错日志
type WriteErr struct {
	Cmd  string
	Name string
}

//日志输出回调
func (wc *WriteLog) Write(p []byte) (int, error) {

	//fmt.Print(string(p) + "---------log---------" + wc.Cmd)
	fmt.Print(string(p))

	//写入日志
	tools.WriteLine("log/"+wc.Name+".log", string(p))

	n := len(p)

	return n, nil
}

//错误输出回调
func (wc *WriteErr) Write(p []byte) (int, error) {

	//fmt.Print(string(p) + "--------err----------" + wc.Cmd)
	fmt.Print(string(p))

	n := len(p)

	return n, nil
}

func RunInit(configLine map[string]string) {

	//fmt.Println(configLine["name"])

	Run(configLine["cmd"], configLine["name"])

}

func Run(cmdLine string, cmdName string) {

	for {

		RunCmd(cmdLine, cmdName)
	}

}

func RunCmd(cmdLine string, cmdName string) {

	//fmt.Println(cmdName)

	var stdoutBuf, stderrBuf bytes.Buffer

	var cmd *exec.Cmd

	sysType := runtime.GOOS

	if sysType == "linux" || sysType == "darwin" {
		// LINUX系统

		cmd = exec.Command("bash", "-c", cmdLine)
	}

	if sysType == "windows" {
		// windows系统

		cmd = exec.Command("cmd", "/c", cmdLine)

	}
	//获取输出
	stdoutIn, _ := cmd.StdoutPipe()
	//获取报错输出
	stderrIn, _ := cmd.StderrPipe()

	defer func() {

		stdoutIn.Close()

		stderrIn.Close()

	}()

	var errStdout, errStderr error

	stdout := io.Writer(&stdoutBuf)
	stderr := io.Writer(&stderrBuf)
	//运行命令
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	//获取命令输出
	go func() {
		counter := &WriteLog{}
		//将当前命令和名称传递过去
		counter.Cmd = cmdLine
		counter.Name = cmdName

		//fmt.Println(cmdName)

		_, errStdout = io.Copy(stdout, io.TeeReader(stdoutIn, counter))

	}()

	//获取命令错误输出
	go func() {

		counter := &WriteErr{}
		//将当前命令和名称传递过去
		counter.Cmd = cmdLine
		counter.Name = cmdName

		_, errStderr = io.Copy(stderr, stderrIn)

	}()

	//等待命令执行
	err = cmd.Wait()

}
