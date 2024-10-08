package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	klogv1 "k8s.io/klog"
	klogv2 "k8s.io/klog/v2"
)

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

	logrus.SetFormatter(&logrus.JSONFormatter{})
	fmt.Println("新的ws已连接。。。", c.Request.RemoteAddr)
	klogv2.Info("V2:新的ws已连接。。。", c.Request.RemoteAddr)
	klogv1.Info("V1:新的ws已连接。。。", c.Request.RemoteAddr)
	logrus.Info("logrus:新的ws已连接。。。", c.Request.RemoteAddr)
	cf := zap.NewProductionConfig()
	zap, _ := cf.Build()
	zap.Info(fmt.Sprintf("zap:新的ws已连接：%s", c.Request.RemoteAddr))
	zap.Info(fmt.Sprintf("zap:新的ws已连接：%s", c.Request.RemoteAddr))
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
