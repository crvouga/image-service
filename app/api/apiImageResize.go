package api

import (
	"fmt"
	"image"
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/library/imageExt"
	"net/http"
	"strconv"
)

// ImageResizeParams defines all available query parameters for the image resize API
type ImageResizeParams struct {
	URL       string `query:"url" doc:"URL of the image to resize" required:"true"`
	Width     int    `query:"width" doc:"Width in pixels (1-2000)" min:"1" max:"2000"`
	Height    int    `query:"height" doc:"Height in pixels (1-2000)" min:"1" max:"2000"`
	ProjectID string `query:"projectID" doc:"Project identifier" required:"true"`
}

func (params *ImageResizeParams) validate() error {
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

// ApiImageResize handles image resizing requests
func ApiImageResize(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		resizedImg, format, err := processImageResize(ac, params)

		if err != nil {
			statusCode := http.StatusInternalServerError
			if err.Error() == "invalid parameters" {
				statusCode = http.StatusBadRequest
			}
			http.Error(w, err.Error(), statusCode)
			return
		}

		fmt.Println("resizedImg", resizedImg)
		fmt.Println("format", format)

		// Set content type based on image format
		// w.Header().Set("Content-Type", "image/"+format)
		// w.Header().Set("Cache-Control", "public, max-age=86400") // Cache for 24 hours

		// // Encode and send the resized image
		// if err := encodeImage(w, resizedImg, format); err != nil {
		// 	http.Error(w, "Failed to encode image: "+err.Error(), http.StatusInternalServerError)
		// 	return
		// }
	}
}

// processImageResize handles the core image resizing logic separate from HTTP concerns
func processImageResize(ac *appCtx.AppCtx, params ImageResizeParams) (image.Image, string, error) {
	if err := params.validate(); err != nil {
		return nil, "", fmt.Errorf("invalid parameters: %w", err)
	}

	projectIDVar, err := projectID.New(params.ProjectID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to parse projectID: %w", err)
	}

	fmt.Println("projectID", projectIDVar)

	project, err := ac.ProjectDB.GetByID(projectIDVar)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get project: %w", err)
	}

	fmt.Println("project", project)

	// Fetch the image
	resp, err := http.Get(params.URL)
	if err != nil {
		return nil, "", fmt.Errorf("failed to fetch image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("failed to fetch image: status %s", resp.Status)
	}

	// Decode the image
	img, format, err := imageExt.Decode(resp.Body, params.URL)
	if err != nil {
		return nil, "", fmt.Errorf("failed to decode image: %w", err)
	}

	fmt.Println("img", img)
	fmt.Println("format", format)

	// Resize the image
	resizedImg := imageExt.Resize(img, params.Width, params.Height)

	return resizedImg, format, nil
}

func parseIntOrZero(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
