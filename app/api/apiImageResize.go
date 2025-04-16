package api

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/library/imageExt"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

// ImageResizeParams defines all available query parameters for the image resize API
type ImageResizeParams struct {
	URL       string `query:"url" doc:"URL of the image to resize" required:"true"`
	Width     int    `query:"width" doc:"Width in pixels (1-2000)" min:"1" max:"2000"`
	Height    int    `query:"height" doc:"Height in pixels (1-2000)" min:"1" max:"2000"`
	ProjectID string `query:"projectID" doc:"Project identifier" required:"true"`
}

// ApiImageResize handles image resizing requests
func ApiImageResize(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Only allow GET requests
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		params := ImageResizeParams{
			URL:       r.URL.Query().Get("url"),
			Width:     parseIntOrZero(r.URL.Query().Get("width")),
			Height:    parseIntOrZero(r.URL.Query().Get("height")),
			ProjectID: r.URL.Query().Get("projectID"),
		}

		if err := validateParams(params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		projectIDVar, err := projectID.New(params.ProjectID)
		if err != nil {
			http.Error(w, "Failed to parse projectID: "+err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("projectID", projectIDVar)

		project, err := ac.ProjectDB.GetByID(projectIDVar)

		if err != nil {
			http.Error(w, "Failed to get project: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("project", project)

		// Fetch the image
		resp, err := http.Get(params.URL)
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
		img, format, err := decodeImage(resp.Body, params.URL)
		if err != nil {
			http.Error(w, "Failed to decode image: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("img", img)
		fmt.Println("format", format)

		imgNew := imageExt.Resize(img, params.Width, params.Height)

		fmt.Println("imgNew", imgNew)

		// Resize the image
		// resizedImg := resize.Resize(uint(params.Width), uint(params.Height), img, resize.Lanczos3)

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

func parseIntOrZero(s string) int {
	i, _ := strconv.Atoi(s)
	return i
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

func validateParams(params ImageResizeParams) error {
	if params.URL == "" {
		return fmt.Errorf("url parameter is required")
	}
	if params.ProjectID == "" {
		return fmt.Errorf("projectID parameter is required")
	}
	if params.Width < 1 || params.Width > 2000 {
		return fmt.Errorf("width must be between 1 and 2000")
	}
	if params.Height < 1 || params.Height > 2000 {
		return fmt.Errorf("height must be between 1 and 2000")
	}
	return nil
}
