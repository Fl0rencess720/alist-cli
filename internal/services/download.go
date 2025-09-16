package services

import (
	"os"

	"github.com/Fl0rencess720/alist-cli/internal/httpclient"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func Download(cmd *cobra.Command, args []string) {
	alistClient := httpclient.GetAlistClient()

	fileName := args[0]

	if err := alistClient.Download(fileName); err != nil {
		zap.L().Error("Download failed", zap.Error(err))
		os.Exit(1)
	}
}
