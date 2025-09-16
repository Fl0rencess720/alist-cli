package cmd

import (
	"github.com/Fl0rencess720/alist-cli/internal/services"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List contents of the current Alist directory",
	Long:    `Lists all files and directories within the current working directory configured for the Alist server.`,
	Example: `alist-cli list`,
	Run:     services.List,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
