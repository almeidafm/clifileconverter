package files

import (
	"fmt"
	"path/filepath"
	"strings"

	_ "golang.org/x/image/webp"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

var textFormats = map[string]struct{}{
	"doc":  {},
	"docx": {},
	"pdf":  {},
}

var imageFormats = map[string]struct{}{
	"jpg":  {},
	"jpeg": {},
	"png":  {},
	"webp": {},
}

var audioFormats = map[string]struct{}{
	"mp3":  {},
	"wav":  {},
	"flac": {},
}

var videoFormats = map[string]struct{}{
	"mp4":  {},
	"mkv":  {},
	"webm": {},
	"mov":  {},
}

func Validate(inputFiles []string) error {
	for _, file := range inputFiles {
		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(file), "."))

		if _, ok := textFormats[ext]; ok {
			continue
		}
		if _, ok := imageFormats[ext]; ok {
			continue
		}
		if _, ok := audioFormats[ext]; ok {
			continue
		}
		if _, ok := videoFormats[ext]; ok {
			continue
		}

		return fmt.Errorf("invalid format: %s", file)
	}

	return nil
}
