package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	GO_VERSION string
	BUILD_TIME string
)
// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", GO_VERSION)
		fmt.Printf("Build Time : %s\n", BUILD_TIME)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
