package cmd

import (
	"github.com/Fl0rencess720/alist-cli/internal/services"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:     "download [filename]",
	Short:   "Download a file from the current Alist directory",
	Long:    `Downloads a single file specified by its name from the current working directory on the Alist server.`,
	Example: `alist-cli download my_document.pdf`,

	Args: cobra.ExactArgs(1),
	Run:  services.Download,
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
