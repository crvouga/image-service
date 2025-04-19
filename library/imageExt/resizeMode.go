package imageExt

type ResizeMode int

const (
	Stretch ResizeMode = iota // Stretch to fill (original behavior)
	Contain                   // Fit within dimensions, adding padding
	Cover                     // Fill dimensions, cropping if needed
)
