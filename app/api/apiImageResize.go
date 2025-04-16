package api

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/library/imageExt"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

// ApiImageResize handles image resizing requests
func ApiImageResize(ac *appContext.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Only allow GET requests
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse query parameters
		imageURL := r.URL.Query().Get("url")
		widthStr := r.URL.Query().Get("width")
		heightStr := r.URL.Query().Get("height")
		projectIDMaybe := r.URL.Query().Get("projectID")

		projectIDInst, err := projectID.New(projectIDMaybe)
		if err != nil {
			http.Error(w, "Failed to parse projectID: "+err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("projectID", projectIDInst)

		project, err := ac.ProjectDB.GetByID(projectIDInst)

		if err != nil {
			http.Error(w, "Failed to get project: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("project", project)

		// Validate URL parameter
		if imageURL == "" {
			http.Error(w, "Missing image URL parameter", http.StatusBadRequest)
			return
		}

		// Parse width and height parameters
		width, height, err := parseResizeDimensions(widthStr, heightStr)
		fmt.Println("width", width)
		fmt.Println("height", height)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Fetch the image
		resp, err := http.Get(imageURL)
		if err != nil {
			http.Error(w, "Failed to fetch image: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Failed to fetch image: status "+resp.Status, http.StatusInternalServerError)
			return
		}

		// Decode the image
		img, format, err := decodeImage(resp.Body, imageURL)
		if err != nil {
			http.Error(w, "Failed to decode image: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("img", img)
		fmt.Println("format", format)

		imgNew := imageExt.Resize(img, width, height)

		fmt.Println("imgNew", imgNew)

		// Resize the image
		// resizedImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

		// // Set content type based on image format
		// w.Header().Set("Content-Type", "image/"+format)
		// w.Header().Set("Cache-Control", "public, max-age=86400") // Cache for 24 hours

		// // Encode and send the resized image
		// if err := encodeImage(w, resizedImg, format); err != nil {
		// 	http.Error(w, "Failed to encode image: "+err.Error(), http.StatusInternalServerError)
		// 	return
		// }
	}
}

// parseResizeDimensions parses and validates width and height parameters
func parseResizeDimensions(widthStr, heightStr string) (int, int, error) {
	var width, height int
	var err error

	// Parse width
	if widthStr != "" {
		width, err = strconv.Atoi(widthStr)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid width parameter: %s", err)
		}
		if width <= 0 || width > 2000 {
			return 0, 0, fmt.Errorf("width must be between 1 and 2000 pixels")
		}
	}

	// Parse height
	if heightStr != "" {
		height, err = strconv.Atoi(heightStr)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid height parameter: %s", err)
		}
		if height <= 0 || height > 2000 {
			return 0, 0, fmt.Errorf("height must be between 1 and 2000 pixels")
		}
	}

	// Ensure at least one dimension is specified
	if width == 0 && height == 0 {
		return 0, 0, fmt.Errorf("at least one of width or height must be specified")
	}

	return width, height, nil

}

// decodeImage decodes an image from the provided reader
func decodeImage(r io.Reader, imageURL string) (image.Image, string, error) {
	// Determine format from file extension
	ext := strings.ToLower(filepath.Ext(imageURL))
	format := strings.TrimPrefix(ext, ".")

	// If format couldn't be determined from URL, try to decode as PNG or JPEG
	if format != "png" && format != "jpeg" && format != "jpg" {
		// Try PNG first
		img, err := png.Decode(r)
		if err == nil {
			return img, "png", nil
		}

		// If PNG fails, try JPEG
		img, err = jpeg.Decode(r)
		if err == nil {
			return img, "jpeg", nil
		}

		return nil, "", fmt.Errorf("unsupported image format")
	}

	// Normalize format
	if format == "jpg" {
		format = "jpeg"
	}

	// Decode based on format
	var img image.Image
	var err error

	switch format {
	case "png":
		img, err = png.Decode(r)
	case "jpeg":
		img, err = jpeg.Decode(r)
	default:
		return nil, "", fmt.Errorf("unsupported image format: %s", format)
	}

	if err != nil {
		return nil, "", err
	}

	return img, format, nil
}

// // encodeImage encodes the image to the provided writer
// func encodeImage(w io.Writer, img image.Image, format string) error {
// 	switch format {
// 	case "png":
// 		return png.Encode(w, img)
// 	case "jpeg":
// 		return jpeg.Encode(w, img, &jpeg.Options{Quality: 85})
// 	default:
// 		return fmt.Errorf("unsupported output format: %s", format)
// 	}
// }
