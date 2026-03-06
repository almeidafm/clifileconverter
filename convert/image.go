package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
	_ "golang.org/x/image/webp"
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
		if err := convertOne(in, toFormat); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("%d file(s) failed: %w", len(errs), errs[0])
	}
	return nil
}

func convertOne(in, toFormat string) error {

	f, err := os.Open(in)
	if err != nil {
		return fmt.Errorf("cannot open %s: %w", in, err)
	}

	img, _, err := image.Decode(f)
	f.Close()
	if err != nil {
		return fmt.Errorf("decode failed for %s: %w", in, err)
	}

	name := strings.TrimSuffix(in, filepath.Ext(in))
	outPath := fmt.Sprintf("%s.%s", name, toFormat)

	if _, err := os.Stat(outPath); err == nil {
		return fmt.Errorf("output file already exists: %s", outPath)
	}

	out, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("cannot create %s: %w", outPath, err)
	}

	switch toFormat {
	case "jpg", "jpeg":
		err = jpeg.Encode(out, img, &jpeg.Options{Quality: 90})
	case "png":
		err = png.Encode(out, img)
	case "webp":
		err = webp.Encode(out, img, &webp.Options{
			Lossless: false,
			Quality:  90,
		})
	}

	out.Close()

	if err != nil {
		os.Remove(outPath)
		return fmt.Errorf("encode failed for %s: %w", in, err)
	}

	fmt.Printf("converted: %s -> %s\n", in, outPath)
	return nil
}
