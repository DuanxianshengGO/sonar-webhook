package route

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"webhook-sonar/model"
	"webhook-sonar/send"
)

func InitRoute(r *gin.Engine) {
	r.POST("/wecom/robot", WeComRobotHandle)
}

func WeComRobotHandle(c *gin.Context) {
	robotHandle(c, send.SendWeComRobotMessage)

}

// 机器人参数处理
func robotHandle(c *gin.Context, sender send.RobotMessageSender) {

	// 1.获取企微机器人key
	var RobotKey string
	c.ShouldBindQuery("key")
	RobotKey = c.DefaultQuery("key", RobotKey)
	//fmt.Println(RobotKey)

	// 2.获取代码质量数据
	data := new(model.WebhookData)
	err := c.BindJSON(data)
	if err != nil {
		log.WithFields(log.Fields{"code": 400, "msg": err.Error()}).Error()
		c.JSON(400, NewResultFail(400, "parse request body error: "+err.Error()))
		return
	}

	//fmt.Printf("dataValue: %#v\n", data)
	log.WithFields(log.Fields{"code": 200, "msg": "get webHookData success, dataValue"}).Info()

	// 3.发送质量报告
	err = sender(RobotKey, data)
	if err != nil {
		c.JSON(500, NewResultFail(1, "request third failed: "+err.Error()))
		return
	}

	c.JSON(200, NewResultOkEmpty())
}

// Result 接口层响应数据
type Result struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 提示消息
	Data interface{} `json:"data"` // 数据
}

func NewResultFail(code int, msg string) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func NewResultOkEmpty() *Result {
	return &Result{
		Code: 0,
		Msg:  "ok",
		Data: nil,
	}
}
