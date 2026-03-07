package cmd

import (
	"fmt"
	"os"

	"github.com/almeidafm/clifileconverter/convert"
	"github.com/almeidafm/clifileconverter/files"
	"github.com/spf13/cobra"
)

var inputFiles []string
var toFormat string

var rootCmd = &cobra.Command{
	Use:   "clifileconvert [file]",
	Short: "Command Line File Converter",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		inputFiles = args

		group, err := files.ValidateFileFormat(inputFiles, toFormat)
		if err != nil {
			return err
		}

		fmt.Println("Group:", group)

		switch group {
		case files.Audio:
			return convert.Audio(inputFiles, toFormat)

		case files.Image:
			return convert.Image(inputFiles, toFormat)

		case files.Video:
			return convert.Video(inputFiles, toFormat)

			return fmt.Errorf("unsupported file group")
		}
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
	rootCmd.Flags().StringVar(&toFormat, "to", "", "target format (jpg, jpeg, png, webp, mp3, wav, flac, mp4, mkv, webm, mov)")
	rootCmd.MarkFlagRequired("to")
}
