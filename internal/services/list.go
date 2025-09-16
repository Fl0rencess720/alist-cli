package services

import (
	"fmt"

	"github.com/Fl0rencess720/alist-cli/internal/httpclient"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func List(cmd *cobra.Command, args []string) {
	alistClient := httpclient.GetAlistClient()

	req, err := alistClient.List(&httpclient.ListReq{
		Path:     viper.GetString("alist_pwd"),
		Password: viper.GetString("ALIST_PASSWORD"),
		Page:     0,
		PerPage:  0,
		Refresh:  true,
	})
	if err != nil {
		zap.L().Error("Failed to list files", zap.Error(err))
		return
	}

	items := req.Data.Content

	dirColor := color.New(color.FgCyan, color.Bold)
	fileColor := color.New(color.FgWhite)
	sizeColor := color.New(color.FgGreen).SprintFunc()
	typeColor := color.New(color.FgYellow).SprintFunc()
	dateColor := color.New(color.FgBlue).SprintFunc()
	for _, item := range items {
		if item.IsDir {
			dirColor.Printf("%-40s", item.Name)
		} else {
			fileColor.Printf("%-40s", item.Name)
		}

		fileTypeStr := typeColor(getFileType(item.IsDir))
		fileSizeStr := sizeColor(formatFileSize(item.Size))
		modifiedStr := dateColor(item.Modified)

		fmt.Printf("  %-10s  %-15s  %s\n", fileTypeStr, fileSizeStr, modifiedStr)
	}

}

func formatFileSize(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}
	s := float64(size)
	i := 0
	for ; i < len(units)-1 && s >= 1024; i++ {
		s /= 1024
	}
	return fmt.Sprintf("%.2f %s", s, units[i])
}

func getFileType(isDir bool) string {
	if isDir {
		return "<DIR>"
	} else {
		return "<FILE>"
	}
}
