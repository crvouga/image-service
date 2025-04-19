package imageExt

import "image"

// HasTransparency checks if an image has any transparent pixels
func HasTransparency(img image.Image) bool {
	if rgba, ok := img.(*image.RGBA); ok {
		bounds := rgba.Bounds()
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				_, _, _, a := rgba.At(x, y).RGBA()
				if a < 0xffff {
					return true
				}
			}
		}
	}
	return false
}
