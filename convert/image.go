package convert

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Image(inputFiles []string, toFormat string) error {
	toFormat = strings.ToLower(toFormat)
	switch toFormat {
	case "jpg", "jpeg", "png", "webp":
	default:
		return fmt.Errorf("unsupported target format: %s", toFormat)
	}

	var errs []error

	for _, in := range inputFiles {
		name := strings.TrimSuffix(in, filepath.Ext(in))
		outPath := fmt.Sprintf("%s.%s", name, toFormat)

		if _, err := os.Stat(outPath); err == nil {
			errs = append(errs, fmt.Errorf("output file already exists: %s", outPath))
			continue
		}

		cmd := exec.Command("ffmpeg", "-y", "-i", in, outPath)
		cmd.Stdout = nil
		cmd.Stderr = nil

		if err := cmd.Run(); err != nil {
			errs = append(errs, fmt.Errorf("conversion failed for %s: %w", in, err))
			continue
		}

		fmt.Printf("converted: %s -> %s\n", in, outPath)
	}

	if len(errs) > 0 {
		return fmt.Errorf("%d file(s) failed: %w", len(errs), errs[0])
	}
	return nil
}
