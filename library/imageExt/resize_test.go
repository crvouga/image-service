package imageExt

import (
	"image"
	"testing"
)

func TestResize(t *testing.T) {
	// Create a test image
	width, height := 100, 80
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	testCases := []struct {
		name         string
		inputImg     image.Image
		targetWidth  int
		targetHeight int
		wantWidth    int
		wantHeight   int
	}{
		{
			name:         "Resize smaller",
			inputImg:     img,
			targetWidth:  50,
			targetHeight: 40,
			wantWidth:    50,
			wantHeight:   40,
		},
		{
			name:         "Resize larger",
			inputImg:     img,
			targetWidth:  200,
			targetHeight: 160,
			wantWidth:    200,
			wantHeight:   160,
		},
		{
			name:         "Same size",
			inputImg:     img,
			targetWidth:  width,
			targetHeight: height,
			wantWidth:    width,
			wantHeight:   height,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resized := Resize(tc.inputImg, tc.targetWidth, tc.targetHeight)

			gotBounds := resized.Bounds()
			gotWidth := gotBounds.Dx()
			gotHeight := gotBounds.Dy()

			if gotWidth != tc.wantWidth {
				t.Errorf("Width = %d, want %d", gotWidth, tc.wantWidth)
			}

			if gotHeight != tc.wantHeight {
				t.Errorf("Height = %d, want %d", gotHeight, tc.wantHeight)
			}
		})
	}
}

func TestResizeWithAlgorithm(t *testing.T) {
	// Create a test image
	width, height := 100, 80
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	targetWidth, targetHeight := 50, 40

	// Test with different algorithms
	algorithms := []ResizeAlgorithm{Nearest, Bilinear}

	for _, algorithm := range algorithms {
		resized := ResizeWithAlgorithm(img, targetWidth, targetHeight, algorithm)

		gotBounds := resized.Bounds()
		gotWidth := gotBounds.Dx()
		gotHeight := gotBounds.Dy()

		if gotWidth != targetWidth {
			t.Errorf("Algorithm %v: Width = %d, want %d", algorithm, gotWidth, targetWidth)
		}

		if gotHeight != targetHeight {
			t.Errorf("Algorithm %v: Height = %d, want %d", algorithm, gotHeight, targetHeight)
		}
	}
}
