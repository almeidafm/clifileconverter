package convert

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func Video(inputFiles []string, toFormat string) error {
	for _, input := range inputFiles {

		ext := filepath.Ext(input)
		name := input[:len(input)-len(ext)]
		output := name + "." + toFormat

		cmd := exec.Command("ffmpeg", "-i", input, output)

		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("conversion failed for %s: %w", input, err)
		}

		fmt.Println("converted:", output)
	}

	return nil
}
