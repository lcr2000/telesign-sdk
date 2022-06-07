package telesign

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/schema"
)

// Client telesign client
type Client struct {
	customerID  string
	apiKey      string
	env         string
	httpTimeOut int
}

// Options optional parameter
type Options struct {
	Env         string
	HttpTimeout int
}

// NewClient Create a new telesign client instance. customerID and apiKey is required
func NewClient(customerID, apiKey string, opts ...*Options) (*Client, error) {
	if customerID == "" || apiKey == "" {
		err := errors.New("customerID and apiKey is required")
		return nil, err
	}
	env := EnvEnterprise
	httpTimeout := DefaultHTTPTimeout
	if len(opts) > 0 {
		opt := opts[0]
		if opt.Env != "" {
			env = opt.Env
		}
		if opt.HttpTimeout > 0 {
			httpTimeout = opt.HttpTimeout
		}
	}
	return &Client{
		customerID:  customerID,
		apiKey:      apiKey,
		env:         env,
		httpTimeOut: httpTimeout,
	}, nil
}

// execute execute the request return a response
func (c *Client) execute(req Requester) ([]byte, error) {
	uri := buildRequestURI(getDomain(c.env), req.GetURI())

	httpReq, err := buildRequest(req.GetMethod(), uri,
		bytes.NewBuffer([]byte(req.GetBody())))
	if err != nil {
		return nil, err
	}

	return execute(c.customerID, c.apiKey, c.httpTimeOut, httpReq,
		req.GetPath(), req.GetBody())
}

// buildRequestURI buildRequestURI
func buildRequestURI(domain string, uri string) string {
	return fmt.Sprintf("%s%s", domain, uri)
}

func buildRequest(method string, uri string, body *bytes.Buffer) (*http.Request, error) {
	httpReq, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}

	if method == http.MethodPost {
		httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	}

	return httpReq, nil
}

func getDomain(env string) string {
	if env == EnvEnterprise {
		return domainEnterprise
	}
	return domain
}

var userAgent = fmt.Sprintf("TeleSignSDK/go-%s Go/%s", version, runtime.Version())

// execute execute the request return a response
func execute(customerID string, apiKey string, httpTimeout int,
	req *http.Request, resource string, body string,
) ([]byte, error) {
	nonce, _ := uuid.NewRandom()
	sigData := buildSignature(time.Now(), nonce.String(), req.Method, resource,
		req.Header.Get("Content-Type"), body)
	sig := fmt.Sprintf("TSA %s:%s", customerID, createSignature(apiKey, sigData))

	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("X-TS-Auth-Method", authMethod)
	req.Header.Add("Authorization", sig)
	req.Header.Add("X-TS-Nonce", sigData.Nonce)
	req.Header.Add("Date", sigData.Date)

	client := getClient(httpTimeout)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

type signatureData struct {
	HTTPMethod  string
	Resource    string
	ContentType string
	Date        string
	Nonce       string
	Body        string
}

func buildSignature(t time.Time, nonce string, method string, resource string,
	contentType string, body string,
) *signatureData {
	return &signatureData{
		HTTPMethod:  method,
		Resource:    resource,
		ContentType: contentType,
		Date:        t.Format(time.RFC1123Z),
		Nonce:       nonce,
		Body:        body,
	}
}

func createSignature(apiKey string, data *signatureData) string {
	var str string
	switch data.HTTPMethod {
	case http.MethodGet, http.MethodDelete:
		str = fmt.Sprintf("%s\n%s\n%s\nx-ts-auth-method:%s\nx-ts-nonce:%s\n%s",
			data.HTTPMethod, data.ContentType, data.Date, authMethod, data.Nonce,
			data.Resource)
	default:
		str = fmt.Sprintf("%s\n%s\n%s\nx-ts-auth-method:%s\nx-ts-nonce:%s\n%s\n%s",
			data.HTTPMethod, data.ContentType, data.Date, authMethod, data.Nonce,
			data.Body, data.Resource)
	}

	key, _ := base64.StdEncoding.DecodeString(apiKey)
	b := hmacSHA256(key, str)
	signature := base64.StdEncoding.EncodeToString(b)

	return signature
}

func hmacSHA256(key []byte, content string) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	return mac.Sum(nil)
}

func getClient(timeout int) http.Client {
	return http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}

// structToURLValues convert struct data to URL Values
func structToURLValues(i interface{}) url.Values {
	val := url.Values{}
	err := schema.NewEncoder().Encode(i, val)
	if err != nil {
		return val
	}
	return val
}
