package imageExt

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Resize resizes an input image to the specified width and height.
// It returns the resized image.
func Resize(img image.Image, width, height int) image.Image {
	// Create a new RGBA image with the target dimensions
	dst := ResizeWithAlgorithm(img, width, height, Bilinear)

	return dst
}

// ResizeWithAlgorithm resizes an image using a specified algorithm.
func ResizeWithAlgorithm(img image.Image, width, height int, algorithm ResizeAlgorithm) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	switch algorithm {
	case Nearest:
		resizeNearest(dst, img)
	case Bilinear:
		resizeBilinear(dst, img)
	default:
		// Default to bilinear
		resizeBilinear(dst, img)
	}

	return dst
}

// resizeNearest implements nearest neighbor interpolation
func resizeNearest(dst *image.RGBA, src image.Image) {
	srcBounds := src.Bounds()
	dstBounds := dst.Bounds()

	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()
	dstW := dstBounds.Dx()
	dstH := dstBounds.Dy()

	for y := range dstH {
		for x := range dstW {
			// Calculate source coordinates
			srcX := int(float64(x) * float64(srcW) / float64(dstW))
			srcY := int(float64(y) * float64(srcH) / float64(dstH))

			// Get color from source and set in destination
			c := src.At(srcBounds.Min.X+srcX, srcBounds.Min.Y+srcY)
			dst.Set(dstBounds.Min.X+x, dstBounds.Min.Y+y, c)
		}
	}
}

// resizeBilinear implements bilinear interpolation
func resizeBilinear(dst *image.RGBA, src image.Image) {
	srcBounds := src.Bounds()
	dstBounds := dst.Bounds()

	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()
	dstW := dstBounds.Dx()
	dstH := dstBounds.Dy()

	// Draw background
	draw.Draw(dst, dstBounds, image.Transparent, image.Point{}, draw.Src)

	for y := range dstH {
		for x := range dstW {
			// Calculate floating point source coordinates
			srcX := float64(x) * float64(srcW-1) / float64(dstW-1)
			srcY := float64(y) * float64(srcH-1) / float64(dstH-1)

			// Calculate the four nearest source pixels
			x0 := int(math.Floor(srcX))
			y0 := int(math.Floor(srcY))
			x1 := int(math.Min(float64(x0+1), float64(srcW-1)))
			y1 := int(math.Min(float64(y0+1), float64(srcH-1)))

			// Calculate interpolation weights
			wx := srcX - float64(x0)
			wy := srcY - float64(y0)

			// Get the four nearest pixels
			c00 := src.At(srcBounds.Min.X+x0, srcBounds.Min.Y+y0)
			c01 := src.At(srcBounds.Min.X+x0, srcBounds.Min.Y+y1)
			c10 := src.At(srcBounds.Min.X+x1, srcBounds.Min.Y+y0)
			c11 := src.At(srcBounds.Min.X+x1, srcBounds.Min.Y+y1)

			// Convert colors to RGBA values
			r00, g00, b00, a00 := c00.RGBA()
			r01, g01, b01, a01 := c01.RGBA()
			r10, g10, b10, a10 := c10.RGBA()
			r11, g11, b11, a11 := c11.RGBA()

			// Bilinear interpolation for each channel
			r := uint8(uint32(bilinearInterpolate(float64(r00), float64(r10), float64(r01), float64(r11), wx, wy)) >> 8)
			g := uint8(uint32(bilinearInterpolate(float64(g00), float64(g10), float64(g01), float64(g11), wx, wy)) >> 8)
			b := uint8(uint32(bilinearInterpolate(float64(b00), float64(b10), float64(b01), float64(b11), wx, wy)) >> 8)
			a := uint8(uint32(bilinearInterpolate(float64(a00), float64(a10), float64(a01), float64(a11), wx, wy)) >> 8)

			// Set the interpolated color
			dst.SetRGBA(dstBounds.Min.X+x, dstBounds.Min.Y+y, color.RGBA{r, g, b, a})
		}
	}
}

// bilinearInterpolate performs bilinear interpolation between four values
func bilinearInterpolate(c00, c10, c01, c11, tx, ty float64) float64 {
	return (1-tx)*(1-ty)*c00 + tx*(1-ty)*c10 + (1-tx)*ty*c01 + tx*ty*c11
}
