package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"webhook-sonar/route"
)

var (
	h        bool
	RobotKey string
	addr     string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&RobotKey, "RobotKey", "", "global wechatrobot webhook, you can overwrite by alert rule with annotations wechatRobot")
	flag.StringVar(&addr, "port", ":9090", "listen addr")
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
		return
	}
	r := gin.Default()
	// 初始化路由
	route.InitRoute(r)
	// 启动服务
	r.Run(addr)
}
