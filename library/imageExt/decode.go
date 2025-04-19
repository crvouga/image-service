package imageExt

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"path/filepath"
	"strings"
)

func Decode(r io.Reader, imageURL string) (image.Image, Format, error) {
	format := getFormatFromURL(imageURL)
	switch format {
	case PNG:
		return decodePNG(r)
	case JPEG:
		return decodeJPEG(r)
	case GIF:
		return decodeGIF(r)
	case BMP:
		return decodeBMP()
	case WEBP:
		return decodeWEBP()
	default:
		return tryDecodeUnknownFormat(r)
	}
}

func decodePNG(r io.Reader) (image.Image, Format, error) {
	img, err := png.Decode(r)
	if err != nil {
		return nil, "", err
	}
	return img, PNG, nil
}

func decodeJPEG(r io.Reader) (image.Image, Format, error) {
	img, err := jpeg.Decode(r)
	if err != nil {
		return nil, "", err
	}
	return img, JPEG, nil
}

func decodeGIF(r io.Reader) (image.Image, Format, error) {
	img, err := gif.Decode(r)
	if err != nil {
		return nil, "", err
	}
	return img, GIF, nil
}

func decodeBMP() (image.Image, Format, error) {
	err := fmt.Errorf("BMP is not supported")
	return nil, "", err
}

func decodeWEBP() (image.Image, Format, error) {
	err := fmt.Errorf("WEBP is not supported")
	return nil, "", err
}

func getFormatFromURL(imageURL string) Format {
	ext := strings.ToLower(filepath.Ext(imageURL))
	ext = strings.TrimPrefix(ext, ".")

	if ext == "jpg" {
		return JPEG
	}

	return Format(ext)
}

func tryDecodeUnknownFormat(r io.Reader) (image.Image, Format, error) {
	// Make a copy of the reader so we can try multiple formats
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, "", err
	}

	// Try each format
	formats := []struct {
		decode func(io.Reader) (image.Image, error)
		format Format
	}{
		{png.Decode, PNG},
		{jpeg.Decode, JPEG},
		{gif.Decode, GIF},
	}

	for _, f := range formats {
		img, err := f.decode(io.NopCloser(strings.NewReader(string(buf))))
		if err == nil {
			return img, f.format, nil
		}
	}

	return nil, "", fmt.Errorf("unsupported image format")
}
