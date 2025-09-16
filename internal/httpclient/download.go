package httpclient

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/viper"
)

type detailReq struct {
	Path     string `json:"path"`
	Password string `json:"password"`
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
	Refresh  bool   `json:"refresh"`
}

type detailResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		IsDir  bool   `json:"isDir"`
		RawURL string `json:"raw_url"`
	} `json:"data"`
}

func (c *Client) GetDetail(detailReq *detailReq) (*detailResp, error) {
	detailResp := detailResp{}

	resp, err := c.client.R().
		SetBody(detailReq).
		SetResult(&detailResp).
		Post(DetailAPIPath)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("http error %d: %s", resp.StatusCode(), resp.String())
	}

	return &detailResp, nil
}

func (c *Client) Download(name string) error {
	filePath := filepath.Join(viper.GetString("alist_pwd"), name)

	detail, err := c.GetDetail(&detailReq{
		Path:     filePath,
		Password: viper.GetString("ALIST_PASSWORD"),
		Page:     1,
		PerPage:  0,
		Refresh:  false,
	})
	if err != nil {
		return err
	}

	if detail.Data.IsDir {
		return fmt.Errorf("'%s' is a directory, cannot download", name)
	}
	if detail.Data.RawURL == "" {
		return fmt.Errorf("could not get a valid download URL for '%s'", name)
	}

	rawURL := detail.Data.RawURL
	outputFileName := name

	headResp, err := c.client.R().Head(rawURL)
	if err != nil {
		return fmt.Errorf("failed to make HEAD request to get file size: %w", err)
	}

	contentLength := headResp.Header().Get("Content-Length")
	if contentLength == "" {
		fmt.Println("Warning: Could not determine file size. Progress bar will not show percentage.")
		contentLength = "-1"
	}

	totalSize, err := strconv.ParseInt(contentLength, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Content-Length header: %w", err)
	}

	resp, err := c.client.R().
		SetDoNotParseResponse(true).
		Get(rawURL)
	if err != nil {
		return fmt.Errorf("download GET request failed: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("http error while downloading %d: received status code %s", resp.StatusCode(), resp.Status())
	}

	body := resp.RawResponse.Body
	defer body.Close()

	destFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %w", outputFileName, err)
	}
	defer destFile.Close()

	bar := progressbar.NewOptions64(
		totalSize,
		progressbar.OptionSetDescription(fmt.Sprintf("Downloading %s", name)),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(30),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	progressReader := progressbar.NewReader(body, bar)
	_, err = io.Copy(destFile, &progressReader)
	if err != nil {
		os.Remove(outputFileName)
		return fmt.Errorf("failed to write to file: %w", err)
	}
	fmt.Printf("✅ Download complete! File saved as '%s'.\n", outputFileName)
	return nil
}
