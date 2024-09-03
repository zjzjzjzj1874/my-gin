package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func main() {
	app := gin.Default()

	app.GET("/hello", func(c *gin.Context) {
		// 输出Hello World
		c.String(200, "Hello World")
	})
	app.GET("/echo", WSEchoHandler)

	if err := app.Run(":9090"); err != nil {
		panic(err)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSEchoHandler(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("新的ws已连接。。。", c.Request.RemoteAddr)
	// 消息处理循环
	for {
		mt, p, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		fmt.Println("message type:", mt)
		fmt.Println("content:", string(p))

		// 回应客户端，这里把消息回写到客户端即可
		_, _ = c.Writer.Write(p)
	}

	fmt.Println("Close WS. Bye~")
	_ = ws.Close()
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"hello": "Gin",
	})
}
