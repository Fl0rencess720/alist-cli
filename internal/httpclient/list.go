package httpclient

import (
	"fmt"
)

type ListReq struct {
	Path     string `json:"path"`
	Password string `json:"password"`
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
	Refresh  bool   `json:"refresh"`
}

type ListResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Content []ListItem `json:"content"`
	} `json:"data"`
}

type ListItem struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	IsDir    bool   `json:"is_dir"`
	Modified string `json:"modified"`
	Sign     string `json:"sign"`
}

func (c *Client) List(listReq *ListReq) (*ListResp, error) {
	listResp := ListResp{}

	resp, err := c.client.R().
		SetBody(listReq).
		SetResult(&listResp).
		Post(ListAPIPath)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("http error %d: %s", resp.StatusCode(), resp.String())
	}

	return &listResp, nil
}
