package telesign

import (
	"encoding/json"
	"log"
	"net/http"
)

// SendSmsVerify Send SMS Verification Code
// Docs: https://developer.telesign.com/enterprise/reference/sendsmsverificationcode
func (c *Client) SendSmsVerify(req *SendSmsVerifyReq) (*SendSmsVerifyResp, error) {
	bytes, err := c.execute(req)
	if err != nil {
		return nil, err
	}
	log.Printf("SendSmsVerify execute complete, uri=%s, body=%s, response=%s\n",
		req.GetURI(), req.GetBody(), string(bytes))
	resp := &SendSmsVerifyResp{}
	err = json.Unmarshal(bytes, resp)
	return resp, err
}

// SendSmsVerifyReq object
type SendSmsVerifyReq struct {
	PhoneNumber   string `schema:"phone_number"`   // required 您要向其发送消息的最终用户的电话号码，以不带空格或特殊字符的数字形式, 以国家/地区拨号代码开头.
	UcID          string `schema:"ucid"`           // 场景 A code specifying the use case you are making the request for.
	OriginatingIP string `schema:"originating_ip"` // 您的最终用户的 IP 地址（不要发送您自己的 IP 地址）. 这用于帮助 TeleSign 改进我们的服务. 支持 IPv4 和 IPv6.
	Language      string `schema:"language"`       // 指定您希望使用的预定义模板的语言的代码. 有关代码的完整列表, 请参阅支持的语言部分. 如果您在 template 参数中提供覆盖消息文本, 则不使用此字段.
	VerifyCode    string `schema:"verify_code"`    // 用于代码质询的验证码. 默认情况下，TeleSign 会为您随机生成一个七位数的数值. 您可以通过在此参数中包含您自己的数字代码来覆盖默认行为, 其值介于000和之间9999999. 无论哪种方式，验证码都会替换消息模板中的变量$$CODE$$.
	Template      string `schema:"template"`       // 覆盖预定义消息模板内容的文本. 包含 $$CODE$$ 变量以自动插入验证码, 最多可包含 1600 个字符.
	SenderID      string `schema:"sender_id"`      // 指定要在 SMS 消息上显示给最终用户的发件人 ID. 在使用它之前, 请将您可能想要使用的任何发件人 ID 提供给我们的客户支持团队, 以便我们可以将它们添加到我们的允许列表中. 如果此字段中的发件人 ID 不在此列表中, 则不使用它. 我们不保证会使用您指定的发件人 ID; TeleSign 可能会覆盖此值以提高交付质量或遵循特定国家/地区的 SMS 法规. TeleSign 建议将值限制为 0-9 和 AZ, 因为对其他 ASCII 字符的支持因运营商而异
	CallbackURL   string `schema:"callback_url"`   // 您希望发送与您的请求相关的交付报告的 URL. 这会覆盖您之前设置的任何默认回调 URL. 覆盖仅持续此请求.
	IsPrimary     string `schema:"is_primary"`     // 无论您是使用此服务作为主要提供者发送此消息 ( ”true”) 还是在主要提供者失败后作为备份 ( ”false”). 我们使用这些数据来优化消息路由.
}

// GetMethod return method request
func (r *SendSmsVerifyReq) GetMethod() string {
	return http.MethodPost
}

// GetURI return uri request
func (r *SendSmsVerifyReq) GetURI() string {
	return r.GetPath()
}

// GetPath return path request
func (r *SendSmsVerifyReq) GetPath() string {
	return SmsVerifyURI
}

// GetBody return body request
func (r *SendSmsVerifyReq) GetBody() string {
	b := structToURLValues(r).Encode()
	return b
}

// SendSmsVerifyResp returned by telesign API
type SendSmsVerifyResp struct {
	MainResponse
	SubResource string  `json:"sub_resource"`
	Errors      []Error `json:"errors"`
	SmsVerify   struct {
		CodeState   string `json:"code_state"`
		CodeEntered string `json:"code_entered"`
	} `json:"verify"`
	ExternalID      string `json:"external_id"`
	SignatureString string `json:"signature_string"`
}
