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

func ValidateFileFormat(inputFiles []string, toFormat string) (FileType, error) {
	if len(inputFiles) == 0 {
		return Unknown, fmt.Errorf("none file provided")
	}

	var group FileType

	for i, file := range inputFiles {
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(file)), ".")
		ft := getFileType(ext)

		if ft == Unknown {
			return Unknown, fmt.Errorf("invalid format: %s", file)
		}

		if i == 0 {
			group = ft
		} else if ft != group {
			return Unknown, fmt.Errorf("different files format not allowed: %s", file)
		}
	}

	toType := getFileType(strings.ToLower(toFormat))
	if toType == Unknown {
		return Unknown, fmt.Errorf("invalid target format: %s", toFormat)
	}

	if group != toType {
		return Unknown, fmt.Errorf("conversion not allowed: %v -> %v", group, toType)
	}

	return group, nil
}
