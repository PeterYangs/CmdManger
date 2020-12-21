package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
	"runtime"
)

//命令输出日志
type WriteLog struct {
	Cmd string
}

//命令报错日志
type WriteErr struct {
	Cmd string
}

func (wc *WriteLog) Write(p []byte) (int, error) {

	fmt.Print(string(p) + "---------log---------" + wc.Cmd)

	n := len(p)

	return n, nil
}

func (wc *WriteErr) Write(p []byte) (int, error) {

	fmt.Print(string(p) + "--------err----------" + wc.Cmd)

	n := len(p)

	return n, nil
}

func Run(cmdLine string) {

	for {

		RunCmd(cmdLine)
	}

}

func RunCmd(cmdLine string) {

	var stdoutBuf, stderrBuf bytes.Buffer

	var cmd *exec.Cmd

	sysType := runtime.GOOS

	if sysType == "linux" {
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
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	go func() {
		counter := &WriteLog{}
		//将当前命令传递过去
		counter.Cmd = cmdLine

		_, errStdout = io.Copy(stdout, io.TeeReader(stdoutIn, counter))

	}()
	go func() {

		counter := &WriteErr{}
		//将当前命令传递过去
		counter.Cmd = cmdLine

		_, errStderr = io.Copy(stderr, stderrIn)

	}()
	err = cmd.Wait()
	//if err != nil {
	//	log.Fatalf("cmd.Run() failed with %s\n", err)
	//}
	//if errStdout != nil || errStderr != nil {
	//	log.Fatal("failed to capture stdout or stderr\n")
	//}
	//outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	//fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

	//stdout.

	//重新拉起
	//RunCmd(cmdLine)

}
