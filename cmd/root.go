package cmd

import (
	"os"

	"github.com/almeidafm/clifileconverter/files"
	"github.com/spf13/cobra"
)

var inputFiles []string

var rootCmd = &cobra.Command{
	Use:   "fileconvert [file]",
	Short: "File Converter",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		inputFiles = args

		return files.Validate(inputFiles)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
