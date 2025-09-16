package httpclient

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

var alistClient *Client

type Client struct {
	client *resty.Client
}

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResp struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func InitClient() {
	baseURL := viper.GetString("ALIST_BASE_URL")
	username := viper.GetString("ALIST_USERNAME")
	password := viper.GetString("ALIST_PASSWORD")

	if baseURL == "" || username == "" || password == "" {
		panic("missing required configuration: ALIST_BASE_URL, ALIST_USERNAME, or ALIST_PASSWORD")
	}

	client := resty.New()
	client.SetBaseURL(baseURL)
	client.SetTimeout(10 * time.Second)

	loginPayload := loginReq{
		Username: username,
		Password: password,
	}
	var loginResult loginResp

	resp, err := client.R().
		SetBody(loginPayload).
		SetResult(&loginResult).
		Post(LoginAPIPath)

	if err != nil {
		panic("failed to send login request: " + err.Error())
	}
	if resp.IsError() {
		panic("login request failed : " + resp.String())
	}
	if loginResult.Code != 200 || loginResult.Data.Token == "" {
		panic("alist login failed : " + loginResult.Message)
	}

	token := loginResult.Data.Token

	client.SetAuthScheme("")
	client.SetAuthToken(token)

	alistClient = &Client{
		client: client,
	}
}

func GetAlistClient() *Client {
	return alistClient
}
