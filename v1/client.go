package binance_connect

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

var BaseURL = [6]string{
	"https://api.binance.com",
	"https://api-gcp.binance.com",
	"https://api1.binance.com",
	"https://api2.binance.com",
	"https://api3.binance.com",
	"https://api4.binance.com",
}


// Endpoint security type
// - If no security type is stated, assume the security type is NONE.
// - API-keys are passed into the Rest API via the X-MBX-APIKEY header.
// - API-keys and secret-keys are case sensitive.
// - API-keys can be configured to only access certain types of secure endpoints.
// 		For example, one API-key could be used for TRADE only,
//  	while another API-key can access everything except for TRADE routes.
// - By default, API-keys can access all secure routes.
type SecurityT int
const (
	None SecurityT = iota // all public access
	Trade // API-key and Singnature required
	UserData // API-key and Singnature required
	UserStream // API-key required
	MARKET_DATA // API-key required
)

type Client struct {
	APIKey     string // API key
	SecretKey  string // Secret key
	BaseURL    string // Base URL for API requests
	HTTPClient *http.Client 
}

// Client factory function
func NewClient(apiKey, secretKey, baseURL string) *Client {
	url := baseURL
	if baseURL == "" {
		url = "https://api.binance.com"
	}
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    url,
		HTTPClient: http.DefaultClient,
	}
}


type Request struct {
	Method   string // http method
	Endpoint string // every api specific url
	SercType SecurityT // security type

	Body     io.Reader 
	Query    url.Values // query string
	Form     url.Values // extually is form data, covert to body in the end
	// header   http.Header
}
func NewBinanceRequest(method, endpoint string, sercType SecurityT) *Request {
	return &Request{
		Method: method,
		Endpoint: endpoint,
		SercType: sercType,

		Query: url.Values{},
		Form: url.Values{},
	}
}

func (r *Request) SetQuery(key string, value interface{}) *Request {
	if r.Query.Get(key) == "" {
		r.Query.Add(key, fmt.Sprintf("%v", value))
		return r
	}
	r.Query.Set(key, fmt.Sprintf("%v", value))
	return r
}
func (r *Request) SetParam(key string, value interface{}) *Request{
	if r.Form.Get(key) == "" {
		r.Form.Add(key, fmt.Sprintf("%v", value))
		return r
	}
	r.Form.Set(key, fmt.Sprintf("%v", value))
	return r
}

type RequsetOption func(*url.Values)

func (c *Client) SetBinanceRequest(r *Request, opts ...RequsetOption) (req *http.Request, err error) {
	if r.SercType == Trade || r.SercType == UserData {
		r.Query.Set("timestamp", fmt.Sprintf("%v",time.Now().UnixNano()/int64(time.Millisecond)))
	}

	bodyString := r.Form.Encode()
	queryString := r.Query.Encode()

	if bodyString != "" {
		r.Body = bytes.NewBufferString(bodyString)
	}
	if r.SercType == Trade || r.SercType == UserData {
		mac := hmac.New(sha256.New, []byte(c.SecretKey))
		_, err = mac.Write([]byte(fmt.Sprintf("%s%s",queryString, bodyString))) // query string + body string 不需要加上 & 符号
		if err != nil {
			return 
		}
		r.Query.Set("signature", fmt.Sprintf("%x", mac.Sum(nil)))
		queryString = r.Query.Encode()
	}

	fullURL := fmt.Sprintf("%s%s?%s", c.BaseURL, r.Endpoint, queryString)

	req, err = http.NewRequest(r.Method, fullURL, r.Body)
	if err != nil {
		return
	}
	log.Printf("full url: %s\nrequese body: %s", req.URL.String(), PrettyPrint(r.Form))

	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s", "binance_connect", "v1"))
	if bodyString != "" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if r.SercType != None {
		req.Header.Set("X-MBX-APIKEY", c.APIKey)
	}
	return 
}


func (c *Client) Call(r *http.Request) (data []byte, err error) {
	resp, err := c.HTTPClient.Do(r)
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}
	defer func() {
		err = resp.Body.Close()
	}()


	if resp.StatusCode != 200 {
		log.Printf("Error: %s", resp.Status)
		return
	}

	// log.Printf("response header: %s", PrettyPrint(resp.Header))

	data, err = io.ReadAll(resp.Body)

	return data, err
}