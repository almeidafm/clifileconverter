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

type FileType int

const (
	Unknown FileType = iota
	Text
	Image
	Audio
	Video
)

func getFileType(ext string) FileType {
	ext = strings.ToLower(ext)
	if _, ok := textFormats[ext]; ok {
		return Text
	}
	if _, ok := imageFormats[ext]; ok {
		return Image
	}
	if _, ok := audioFormats[ext]; ok {
		return Audio
	}
	if _, ok := videoFormats[ext]; ok {
		return Video
	}
	return Unknown
}

func ValidateFileFormat(inputFiles []string, toFormat string) error {
	if len(inputFiles) == 0 {
		return fmt.Errorf("none file provided")
	}

	var group FileType

	for i, file := range inputFiles {
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(file)), ".")
		ft := getFileType(ext)

		if ft == Unknown {
			return fmt.Errorf("invalid format: %s", file)
		}

		if i == 0 {
			group = ft
		} else if ft != group {
			return fmt.Errorf("different files format not allowed: %s", file)
		}
	}

	toType := getFileType(strings.ToLower(toFormat))

	if toType == Unknown {
		return fmt.Errorf("invalid target format: %s", toFormat)
	}

	if toType != group {
		return fmt.Errorf("cannot convert %v files to %s", group, toFormat)
	}

	return nil
}
