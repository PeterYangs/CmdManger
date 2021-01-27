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

		context.String(200, "123")
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
		global.GlobalStatus.CmdList = append(global.GlobalStatus.CmdList, map[string]string{"name": value["name"], "cmd": value["cmd"], "num": value["num"]})
		global.GlobalLock.Unlock()

		//启动命令
		for i := 0; i < num; i++ {

			go func(temp map[string]string) {

				cmd.RunInit(temp)

			}(value)

		}

	}

}
