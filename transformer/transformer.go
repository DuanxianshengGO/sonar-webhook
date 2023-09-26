package transformer

import (
	"bytes"
	"fmt"
	"time"
	"webhook-sonar/model"
)

func TransformToMarkDown(webhookData *model.WebhookData) (markdown *model.WechatMarkdown) {
	var buffer bytes.Buffer

	switch webhookData.QualityGate.Status {
	case "OK":
		buffer.WriteString(fmt.Sprintf("### 代码检测结果: <font color=\"info\">%s</font> \n", "通过"))
	case "ERROR":
		buffer.WriteString(fmt.Sprintf("### 代码检测结果: <font color=\"warning\">`%s`</font> \n", "未通过"))
	}

	buffer.WriteString(fmt.Sprintf("\n>CheckRule: %s\n", webhookData.QualityGate.Name))
	buffer.WriteString(fmt.Sprintf("\n>Project: %s\n", webhookData.Project.Name))
	buffer.WriteString(fmt.Sprintf("\n>Branch: %s\n", webhookData.Branch.Name))

	for _, contion := range webhookData.QualityGate.Conditions {
		if contion.Metric == "bugs" {
			if contion.Status == "ERROR" {
				buffer.WriteString(fmt.Sprintf("\n>Bugs: <font color=\"warning\">`%s`</font>\n", contion.Value))
			} else {
				buffer.WriteString(fmt.Sprintf("\n>Bugs: <font color=\"info\">%s</font>\n", contion.Value))
			}

		}
		if contion.Metric == "code_smells" {
			if contion.Status == "ERROR" {
				buffer.WriteString(fmt.Sprintf("\n>异味: <font color=\"warning\">`%s`</font>\n", contion.Value))
			} else {
				buffer.WriteString(fmt.Sprintf("\n>异味: <font color=\"info\">%s</font>\n", contion.Value))
			}

		}
		if contion.Metric == "vulnerabilities" {
			if contion.Status == "ERROR" {
				buffer.WriteString(fmt.Sprintf("\n>漏洞: <font color=\"warning\">`%s`</font>\n", contion.Value))
			} else {
				buffer.WriteString(fmt.Sprintf("\n>漏洞: <font color=\"info\">%s</font>\n", contion.Value))
			}

		}
		if contion.Metric == "coverage" {
			if contion.Status == "ERROR" {
				buffer.WriteString(fmt.Sprintf("\n>覆盖率: <font color=\"warning\">`%s%%`</font>\n", contion.Value))
			} else {
				buffer.WriteString(fmt.Sprintf("\n>覆盖率: <font color=\"info\">%s%%</font>\n", contion.Value))
			}

		}
		if contion.Metric == "duplicated_lines_density" {
			if contion.Status == "ERROR" {
				buffer.WriteString(fmt.Sprintf("\n>重复率: <font color=\"warning\">`%s%%`</font>\n", contion.Value))
			} else {
				buffer.WriteString(fmt.Sprintf("\n>重复率: <font color=\"info\">%s%%</font>\n", contion.Value))
			}

		}

		if contion.Metric != "duplicated_lines_density" && contion.Metric != "coverage" && contion.Metric != "vulnerabilities" &&
			contion.Metric != "code_smells" && contion.Metric != "bugs" {
			if contion.Status == "ERROR" {
				buffer.WriteString(fmt.Sprintf("\n>%s: <font color=\"warning\">`%s`</font>\n", contion.Metric, contion.Value))
			} else if contion.Status == "NO_VALUE" {
				continue
			} else {
				buffer.WriteString(fmt.Sprintf("\n>%s: <font color=\"info\">%s</font>\n", contion.Metric, contion.Value))
			}
		}

	}
	buffer.WriteString(fmt.Sprintf("\n>触发时间: %s\n", time.Now().Format("2006-01-02 15:04:05")))
	buffer.WriteString(fmt.Sprintf("\n>报告详情: [点击进入](%s)\n", webhookData.Branch.URL))

	markdown = &model.WechatMarkdown{
		MsgType: "markdown",
		MarkDown: &model.Markdown{
			Content: buffer.String(),
		},
	}
	return
}
