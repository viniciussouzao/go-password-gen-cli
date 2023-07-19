package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciussouzao/go-password-gen-cli/pkg/beauty"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the version of the cli",
	Long: `Show the version of the cli
Usage: 
 $ password-gen version

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(beauty.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
