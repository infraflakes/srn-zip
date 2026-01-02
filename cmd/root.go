package cmd

import (
	"github.com/spf13/cobra"
	"github.com/infraflakes/srn-zip/pkg"
)

var RootCmd = &cobra.Command{
	Use:   "szip",
	Short: "Archive and extract files with 7z",
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(pkg.ZipCmd)
	RootCmd.AddCommand(pkg.UnzipCmd)
}
