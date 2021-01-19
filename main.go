package main

import (
	"cmdManger/cmd"
	_ "cmdManger/cmd"
	"encoding/json"
	"github.com/PeterYangs/tools"
	"log"
	"strconv"
	_ "time"
)

func main() {

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

		for i := 0; i < num; i++ {

			go func() {

				cmd.Run(value["cmd"])

			}()

		}

	}

	for {

	}

	//fmt.Println(jsons.List[0]["name"])

	//var _ []map[string]string

	//fmt.Println(jsons["list"])

	//list:=jsons["list"]
	//
	//lists:=list.([]map[string]interface{})
	//
	//
	//fmt.Println(lists)

	//list :=jsons["list"].([]map[string]string)
	//
	//fmt.Println(list)

	//fmt.Println(jsons["list"][0])

	//for i := 0; i < 10; i++ {
	//
	//	go func() {
	//
	//		cmd.Run("php index.php")
	//
	//	}()
	//
	//}
	//
	//time.Sleep(10 * time.Hour)

	//r := gin.Default()
	//
	//r.Static("/static", "./static")
	//
	//r.LoadHTMLGlob("templates/*")
	//
	//r.GET("/status", func(c *gin.Context) {
	//
	//	c.HTML(http.StatusOK, "status.html", gin.H{})
	//
	//})
	//
	//r.Run()

}
