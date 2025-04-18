package imageExt

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"path/filepath"
	"strings"
)

func Decode(r io.Reader, imageURL string) (image.Image, string, error) {
	// Get format from file extension
	ext := strings.ToLower(filepath.Ext(imageURL))
	format := strings.TrimPrefix(ext, ".")

	// Normalize jpg to jpeg
	if format == "jpg" {
		format = "jpeg"
	}

	// Try to decode based on format
	var img image.Image
	var err error

	switch format {
	case "png":
		img, err = png.Decode(r)
	case "jpeg":
		img, err = jpeg.Decode(r)
	default:
		// Try PNG first for unknown formats
		img, err = png.Decode(r)
		if err == nil {
			return img, "png", nil
		}

		// Then try JPEG
		img, err = jpeg.Decode(r)
		if err == nil {
			return img, "jpeg", nil
		}

		return nil, "", fmt.Errorf("unsupported image format")
	}

	if err != nil {
		return nil, "", err
	}

	return img, format, nil
}
