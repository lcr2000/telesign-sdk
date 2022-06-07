package telesign

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// SendSms Send an SMS message
// Docs: https://developer.telesign.com/enterprise/reference/post_v1-messaging
func (c *Client) SendSms(req *SendSmsReq) (*SendSmsResp, error) {
	// convert the spaces into "%20" first, later we replace all "%2520" into "%20"
	t := url.URL{Path: req.Message}
	req.Message = t.String()
	bytes, err := c.execute(req)
	if err != nil {
		return nil, err
	}
	log.Printf("SendSms execute complete, uri=%s, body=%s, response=%s\n",
		req.GetURI(), req.GetBody(), string(bytes))
	resp := &SendSmsResp{}
	err = json.Unmarshal(bytes, resp)
	return resp, err
}

// SendSmsReq object
type SendSmsReq struct {
	PhoneNumber           string `schema:"phone_number"`            // required 包含国家/地区代码的最终用户电话号码. 避免使用特殊字符和空格.
	Message               string `schema:"message"`                 // required 要发送给最终用户的消息文本. 您被限制为 1600 个字符. 如果您发送很长的消息, TeleSign 会将您的消息拆分为单独的部分. TeleSign 建议尽可能不要发送需要多条 SMS 的消息.
	MessageType           string `schema:"message_type"`            // required 此参数指定消息中发送的流量类型. 您可以提供以下值之一: OTP（一次性密码） ARN（警报、提醒和通知） MKT（营销流量）.
	AccountLifecycleEvent string `schema:"account_lifecycle_event"` // 此参数允许您指示您在发送交易时处于生命周期的哪个阶段. 此参数的选项是 - create - 用于创建新帐户 登录 最终用户登录其帐户时. 交易 - 当最终用户在其帐户中完成交易时 update - 执行更新时, 例如更新帐户信息或类似信息. 删除 - 删除帐户时.
	SenderID              string `schema:"sender_id"`               // 指定要在 SMS 消息上显示给最终用户的发件人 ID. 在使用它之前, 请将您可能想要使用的任何发件人 ID 提供给我们的客户支持团队, 以便我们可以将它们添加到我们的允许列表中. 如果此字段中的发件人 ID 不在此列表中, 则不使用它. 我们不保证会使用您指定的发件人 ID; TeleSign 可能会覆盖此值以提高交付质量或遵循特定国家/地区的 SMS 法规。我们建议将值限制为 0-9 和 AZ, 因为对其他 ASCII 字符的支持因运营商而异.
	ExternalID            string `schema:"external_id"`             // 客户为此交易生成的 ID. 响应只是回显为此参数提供的值.
	OriginatingIP         string `schema:"originating_ip"`          // 您的最终用户的 IP 地址（不要发送您自己的 IP 地址）. 此值必须采用 Internet 工程任务组 (IETF) 在标题为 IPv4 和 IPv6 地址的文本表示的 Internet 草案文档中定义的格式.
	CallbackURL           string `schema:"callback_url"`            // 您希望发送与您的请求相关的交付报告的 URL. 这会覆盖您之前设置的任何默认回调 URL. 覆盖仅持续此请求.
	IsPrimary             string `schema:"is_primary"`              // 无论您是使用此服务作为主要提供者发送此消息 ( ”true”) 还是在主要提供者失败后作为备份 ( ”false”). 我们使用这些数据来优化消息路由.
}

// GetMethod return method request
func (r *SendSmsReq) GetMethod() string {
	return http.MethodPost
}

// GetURI return uri request
func (r *SendSmsReq) GetURI() string {
	return r.GetPath()
}

// GetPath return path request
func (r *SendSmsReq) GetPath() string {
	return SmsURI
}

// GetBody return body request
func (r *SendSmsReq) GetBody() string {
	b := structToURLValues(r).Encode()
	b = strings.ReplaceAll(b, "%2520", "%20")
	return b
}

// SendSmsResp returned by telesign API
type SendSmsResp struct {
	MainResponse
	AdditionalInfo AdditionalInfo `json:"additional_info"`
}
