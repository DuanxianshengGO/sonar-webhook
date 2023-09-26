package send

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"webhook-sonar/model"
	"webhook-sonar/transformer"
)

type RobotMessageSender func(robotKey string, webhookData *model.WebhookData) error

func SendWeComRobotMessage(robotKey string, webhookData *model.WebhookData) (err error) {
	markdown := transformer.TransformToMarkDown(webhookData)
	data, err := json.Marshal(markdown)
	if err != nil {
		log.WithFields(log.Fields{"code": 400, "msg": err.Error()}).Error()
		return
	}
	weChatRobotUrl := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + robotKey

	// 发送请求到微信
	req, err := http.NewRequest(
		"POST",
		weChatRobotUrl,
		bytes.NewBuffer(data),
	)
	if err != nil {
		log.WithFields(log.Fields{"code": 400, "msg": err.Error()}).Error()
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"code": 400, "msg": err.Error()}).Error()
		return
	}

	defer resp.Body.Close()

	return

}
