package model

const (
	StatusSuccess = "SUCCESS"
	StatusError   = "ERROR"
	StatusOk      = "OK"
)

// RobotWebhookParameters 机器人回掉请求参数
type RobotWebhookParameters struct {
	Key         string `json:"key" form:"key" binding:"required"` // 访问密钥
	SonarToken  string `json:"sonarToken" form:"sonarToken"`      // SonarQue认证密钥
	SkipSuccess bool   `json:"skipSuccess" form:"skipSuccess"`    // 是否跳过质量门成功
	NewCode     bool   `json:"newCode" form:"newCode"`            // 是否显示新代码指标
}

// WebhookData webhook请求参数
type WebhookData struct {
	ServerURL   string      `json:"serverUrl"`   // 服务地址
	TaskID      string      `json:"taskId"`      // 任务Id
	Status      string      `json:"status"`      // 任务状态
	AnalysedAt  string      `json:"analysedAt"`  // 分析时间
	Revision    string      `json:"revision"`    // 修订号
	Project     Project     `json:"project"`     // 项目信息
	Branch      Branch      `json:"branch"`      // 分支信息
	QualityGate QualityGate `json:"qualityGate"` // 质量阈值信息
	Properties  Properties  `json:"properties"`  // 附加属性
}

// Project 项目名称
type Project struct {
	Key  string `json:"key"`  // 项目标识
	Name string `json:"name"` // 项目名称
	URL  string `json:"url"`  // 项目访问url
}

type Branch struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	IsMain bool   `json:"isMain"`
	URL    string `json:"url"`
}

type QualityGate struct {
	Name       string       `json:"name"`       // 质量配置名称
	Status     string       `json:"status"`     // 质量阈值是否通过
	Conditions []Conditions `json:"conditions"` // 质量配置条件
}

type Conditions struct {
	Metric         string `json:"metric"`         // 指标名称
	Operator       string `json:"operator"`       // 操作符
	ErrorThreshold string `json:"errorThreshold"` // 阈值
	Value          string `json:"value"`          // 指标值
	Status         string `json:"status"`         // 状态
	OnLeakPeriod   bool   `json:"OnLeakPeriod"`
}

type Properties struct {
	DetectedScm string `json:"sonar.analysis.detectedscm"`
	DetectedCI  string `json:"sonar.analysis.detectedci"`
}

// IsQualityGateSuccess 判断代码质量是否过关
func (w WebhookData) IsQualityGateSuccess() bool {
	return w.QualityGate.Status == StatusOk || w.QualityGate.Status == StatusSuccess
}

type Markdown struct {
	Content string `json:"content"`
}

type WechatMarkdown struct {
	MsgType  string    `json:"msgtype"`
	MarkDown *Markdown `json:"markdown"`
}
