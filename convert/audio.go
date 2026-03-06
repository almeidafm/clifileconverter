package convert

import (
	"fmt"
	"path/filepath"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func Audio(inputFiles []string, toFormat string) error {

	toFormat = strings.ToLower(toFormat)

	for _, in := range inputFiles {

		ext := filepath.Ext(in)
		name := strings.TrimSuffix(in, ext)

		out := fmt.Sprintf("%s.%s", name, toFormat)

		err := ffmpeg.
			Input(in).
			Output(out).
			OverWriteOutput().
			Run()

		if err != nil {
			return fmt.Errorf("conversion failed for %s: %w", in, err)
		}

		fmt.Printf("converted: %s -> %s\n", in, out)
	}

	return nil
}
