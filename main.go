package main

import (
	_ "cmdManger/cmd"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "time"
)

func main() {

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

	r := gin.Default()

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/*")

	r.GET("/status", func(c *gin.Context) {

		c.HTML(http.StatusOK, "status.html", gin.H{})

	})

	r.Run()

}
