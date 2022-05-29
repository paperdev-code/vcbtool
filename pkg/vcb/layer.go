// Layers are a structure containing the
package vcb

import "fmt"

type AnyPixel interface {
	ToColor() Color
	ToInk() Ink
}

type AnyLayer interface {
	FromImage(Image) error
	ToImage() (Image, error)
	Set(AnyPixel, int, int)
	Get(int, int) AnyPixel
	GetDimensions() (int, int)
	Clear()
}

// Logical layer
// Contains the 'Inks' that builds up the circuit
type LogicLayer struct {
	Width, Height int
	Circuit       []Ink
}

// Decoration On Layer
// Contains the RGBA color data
type DecorOnLayer struct {
	Width, Height int
	Pixels        []Color
}

// Decoration Off Layer
// Contains the RGBA color data
type DecorOffLayer struct {
	Width, Height int
	Pixels        []Color
}

// Create a new layer
func NewLayer[T interface {
	LogicLayer | DecorOnLayer | DecorOffLayer
}](width int, height int) (*T, error) {
	layer := new(T)

	switch l := any(layer).(type) {
	case *LogicLayer:
		l.Width, l.Height = width, height
		l.Circuit = make([]Ink, width*height)

	case *DecorOnLayer:
		l.Width, l.Height = width, height
		l.Pixels = make([]Color, width*height)

	case *DecorOffLayer:
		l.Width, l.Height = width, height
		l.Pixels = make([]Color, width*height)

	default:
		return nil, fmt.Errorf("invalid type '%T'", l)
	}

	return layer, nil
}

// Clear the logic layer
func (l *LogicLayer) Clear() {
	for i := 0; i < len(l.Circuit); i++ {
		l.Circuit[i] = InkNone
	}
}

// Clear the decoration (On) layer
func (l *DecorOnLayer) Clear() {
	for i := 0; i < len(l.Pixels); i++ {
		l.Pixels[i] = Color{0, 0, 0, 0}
	}
}

// Clear the decoration (Off) image
func (l *DecorOffLayer) Clear() {
	for i := 0; i < len(l.Pixels); i++ {
		l.Pixels[i] = Color{0, 0, 0, 0}
	}
}

// Get ink inside data
func (l *LogicLayer) Get(x int, y int) AnyPixel {
	if (x < 0 || x >= l.Width) || (y < 0 || y >= l.Height) {
		return InkNone
	}
	i := y*l.Width + x
	return l.Circuit[i]
}

// Get color inside data
func (l *DecorOnLayer) Get(x int, y int) AnyPixel {
	if (x < 0 || x >= l.Width) || (y < 0 || y >= l.Height) {
		return Color{0, 0, 0, 0}
	}
	i := y*l.Width + x
	return l.Pixels[i]
}

// Get color inside data
func (l *DecorOffLayer) Get(x int, y int) AnyPixel {
	if (x < 0 || x >= l.Width) || (y < 0 || y >= l.Height) {
		return Color{0, 0, 0, 0}
	}
	i := y*l.Width + x
	return l.Pixels[i]
}

// Set ink inside data
func (l *LogicLayer) Set(pixel AnyPixel, x int, y int) {
	if (x < 0 || x >= l.Width) || (y < 0 || y >= l.Height) {
		return
	}
	i := y*l.Width + x
	l.Circuit[i] = pixel.ToInk()
}

// Set color inside data
func (l *DecorOnLayer) Set(pixel AnyPixel, x int, y int) {
	if (x < 0 || x >= l.Width) || (y < 0 || y >= l.Height) {
		return
	}
	i := y*l.Width + x
	l.Pixels[i] = pixel.ToColor()
}

// Set color inside data
func (l *DecorOffLayer) Set(pixel AnyPixel, x int, y int) {
	if (x < 0 || x >= l.Width) || (y < 0 || y >= l.Height) {
		return
	}
	i := y*l.Width + x
	l.Pixels[i] = pixel.ToColor()
}

// Convert logic layer to colors
func (l *LogicLayer) ToImage() (Image, error) {
	colors := make([]Color, len(l.Circuit))
	for i, ink := range l.Circuit {
		colors[i] = ink.ToColor()
	}
	return colors, nil
}

// Convert decoration layer (On) to image
func (l *DecorOnLayer) ToImage() (Image, error) {
	image := make([]Color, len(l.Pixels))
	copy(image, l.Pixels)
	return image, nil
}

// Convert decoration later (Off) to image
func (l *DecorOffLayer) ToImage() (Image, error) {
	image := make([]Color, len(l.Pixels))
	copy(image, l.Pixels)
	return image, nil
}

// Populate logic layer from colors
func (l *LogicLayer) FromImage(colors Image) error {
	size := l.Width * l.Height
	if size < 0 || size > len(colors) {
		return fmt.Errorf("layer (%d) and colors (%d) size mismatch", size, len(colors))
	}
	for i, color := range colors {
		l.Circuit[i] = color.ToInk()
	}
	return nil
}

// Populate a decoration layer (On) from colors
func (l *DecorOnLayer) FromImage(colors Image) error {
	size := l.Width * l.Height
	if size < 0 || size > len(colors) {
		return fmt.Errorf("layer (%d) and colors (%d) size mismatch", size, len(colors))
	}
	copy(l.Pixels, colors)
	return nil
}

// Populate a decoration layer (Off) from colors
func (l *DecorOffLayer) FromImage(colors Image) error {
	size := l.Width * l.Height
	if size < 0 || size > len(colors) {
		return fmt.Errorf("layer (%d) and colors (%d) size mismatch", size, len(colors))
	}
	copy(l.Pixels, colors)
	return nil
}

// Get width and height of layer
func (l *LogicLayer) GetDimensions() (int, int) {
	return l.Width, l.Height
}

// Get width and height of layer
func (l *DecorOnLayer) GetDimensions() (int, int) {
	return l.Width, l.Height
}

// Get width and height of layer
func (l *DecorOffLayer) GetDimensions() (int, int) {
	return l.Width, l.Height
}
