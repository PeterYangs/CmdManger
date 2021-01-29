package cmd

import (
	"bytes"
	"cmdManger/global"
	"context"
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

		item := global.GetCmdListByName(cmdName)

		temp := *item

		if temp["status"] != global.Success {

			return
		}

		RunCmd(cmdLine, cmdName)
	}

}

func RunCmd(cmdLine string, cmdName string) {

	//fmt.Println(cmdName)

	var stdoutBuf, stderrBuf bytes.Buffer

	var cmd *exec.Cmd

	sysType := runtime.GOOS

	ctx, cancel := context.WithCancel(context.Background())

	if sysType == "linux" || sysType == "darwin" {
		// linux/mac系统

		cmd = exec.CommandContext(ctx, "bash", "-c", cmdLine)
	}

	if sysType == "windows" {
		// windows系统

		cmd = exec.CommandContext(ctx, "cmd", "/c", cmdLine)

	}

	global.GlobalLock.Lock()

	tempMap := make(map[string][]*context.CancelFunc)

	if len(global.GlobalStatus.CancelFuncList) != 0 {

		tempMap = global.GlobalStatus.CancelFuncList
	}

	tempMap[cmdName] = append(tempMap[cmdName], &cancel)

	global.GlobalStatus.CancelFuncList = tempMap

	global.GlobalLock.Unlock()

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

	//go func(cancelFunc context.CancelFunc) {
	//
	//	time.Sleep(time.Second*5)
	//
	//	cancelFunc()
	//
	//}(cancel)

	//等待命令执行
	err = cmd.Wait()

}
