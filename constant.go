package telesign

// internal constant
const (
	version          = "1.0.0"
	authMethod       = "HMAC-SHA256"
	domain           = "https://rest-api.telesign.com"
	domainEnterprise = "https://rest-ww.telesign.com"
)

const (
	// DefaultHTTPTimeout Default http interface timeout
	DefaultHTTPTimeout = 10
)

// API ENV
const (
	// EnvStandard is the standard env
	EnvStandard = "Standard"
	// EnvEnterprise is the enterprise env
	EnvEnterprise = "Enterprise"
)

// URI
const (
	SmsVerifyURI = "/v1/verify/sms"
	SmsURI       = "/v1/messaging"
)

// MessageType
const (
	// MessageTypeARN type
	MessageTypeARN = "ARN"
	// MessageTypeMKT type
	MessageTypeMKT = "MKT"
	// MessageTypeOTP type
	MessageTypeOTP = "OTP"
)

// IsPrimary
const (
	IsPrimaryTrue  = "true"
	IsPrimaryFalse = "false"
)

// Error Code
// Docs: https://developer.telesign.com/enterprise/docs/all-status-and-error-codes
const (
	// ErrorCodeAccountLimitReached 已达到账号限额
	ErrorCodeAccountLimitReached = -30003
)
