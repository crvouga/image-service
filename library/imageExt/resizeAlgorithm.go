package imageExt

// ResizeAlgorithm represents the algorithm to use for image resizing
type ResizeAlgorithm string

const (
	// Nearest uses nearest neighbor interpolation
	Nearest ResizeAlgorithm = "nearest"
	// Bilinear uses bilinear interpolation
	Bilinear ResizeAlgorithm = "bilinear"
)
