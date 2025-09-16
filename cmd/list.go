package cmd

import (
	"github.com/Fl0rencess720/alist-cli/internal/services"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列举当前所处Alist目录下的所有文件",
	Long:  `列举当前所处Alist目录下的所有文件`,
	Run:   services.List,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
