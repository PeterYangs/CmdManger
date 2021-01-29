package main

import (
	"cmdManger/cmd"
	_ "cmdManger/cmd"
	"cmdManger/global"
	"encoding/json"
	"fmt"
	"github.com/PeterYangs/tools"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	_ "time"
)

func main() {

	//启动脚本
	go runCmd()

	r := gin.Default()

	r.Delims("{[{", "}]}")

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/*")

	r.GET("/status", func(c *gin.Context) {

		c.HTML(http.StatusOK, "status.html", gin.H{})

	})

	r.GET("/getStatus", func(context *gin.Context) {

		context.JSON(200, global.GlobalStatus.CmdList)
	})

	r.GET("/printCancel", func(context *gin.Context) {

		fmt.Println(global.GlobalStatus.CancelFuncList)

	})

	//取消指定名称的进程
	r.GET("/stopProcess", func(context *gin.Context) {

		name, bools := context.GetQuery("name")

		if bools != true {

			context.JSON(200, gin.H{"code": 2, "msg": "no name"})

			return

		}

		cmdItem := global.GetCmdListByName(name)

		cmdItemTemp := *cmdItem

		if cmdItemTemp == nil {

			context.JSON(200, gin.H{"code": 2, "msg": "no match cmd"})

			return
		}

		cmdItemTemp["status"] = global.Stop

		cancelList := global.GlobalStatus.CancelFuncList[name]

		for _, cancel := range cancelList {

			c := *cancel

			c()

		}

		context.JSON(200, gin.H{"code": 1, "msg": "success", "data": name})

	})

	r.Run()

}

func runCmd() {

	data, err := tools.ReadFile("./config.json")

	if err != nil {

		log.Fatal(err)

	}

	type config struct {
		List []map[string]string
	}

	var jsons config

	err = json.Unmarshal([]byte(data), &jsons)

	if err != nil {

		log.Fatal(err)
	}

	for _, value := range jsons.List {

		num, err := strconv.Atoi(value["num"])

		if err != nil {

			log.Fatal(err)

		}

		if num <= 0 {

			panic("启动数量不能小于等于0")
		}

		global.GlobalLock.Lock()
		global.GlobalStatus.CmdList = append(global.GlobalStatus.CmdList, map[string]string{"name": value["name"], "cmd": value["cmd"], "num": value["num"], "status": global.Success})
		global.GlobalLock.Unlock()

		//启动命令
		for i := 0; i < num; i++ {

			go func(temp map[string]string) {

				cmd.RunInit(temp)

			}(value)

		}

	}

}
