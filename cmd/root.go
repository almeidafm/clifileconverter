package cmd

import (
	"fmt"
	"os"

	"github.com/almeidafm/clifileconverter/files"
	"github.com/spf13/cobra"
)

var inputFiles []string
var toFormat string

var rootCmd = &cobra.Command{
	Use:   "fileconvert [file]",
	Short: "File Converter",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		inputFiles = args

		if err := files.ValidateFileFormat(inputFiles); err != nil {
			return err
		}

		fmt.Println("Sucess!!")

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&toFormat, "to", "", "target format (doc, docx, pdf, jpg, jpeg, png, webp, mp3, wav, flac, mp4, mkv, webm, mov)")
	rootCmd.MarkFlagRequired("to")
}
