# telesign-sdk
telesign Golang SDK

## quick start
```go
import (
	"github.com/lcr2000/telesign-sdk"
	"log"
)

func main() {
	client, err := telesign.NewClient("your customerID", "your apiKey")
	if err != nil {
		log.Printf("NewClient fail, err=%v", err)
		return
	}
	req := &telesign.SendSmsVerifyReq{
		PhoneNumber:   "",
		UcID:          "",
		OriginatingIP: "",
		Language:      "",
		VerifyCode:    "",
		Template:      "",
		SenderID:      "",
		CallbackURL:   "",
		IsPrimary:     "",
	}
	resp, err := client.SendSmsVerify(req)
	if err != nil || resp == nil {
		log.Printf("SendSmsVerify fail, resp=%+v, err=%v", resp, err)
		return
	}
	log.Printf("SendSmsVerify resp=%+v", resp)
}
```
