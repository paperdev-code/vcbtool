// Virtual Circuit Board Package
// Functions for parsing the image inside of blueprints
package vcb

import (
	"fmt"
)

type RGBA = struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

// Ink is any of the possible usable logical inks
type Ink int8

// Ink enum
const (
	INK_NONE Ink = iota

	INK_WRITE
	INK_CROSS
	INK_READ
	INK_BUFFER
	INK_AND
	INK_OR
	INK_XOR
	INK_NOT
	INK_NAND
	INK_NOR
	INK_XNOR
	INK_LATCH_ON
	INK_LATCH_OFF
	INK_CLOCK
	INK_LED
	INK_ANNOTATION
	INK_FILLER

	INK_TC_GRAY
	INK_TC_WHITE
	INK_TC_RED
	INK_TC_ORANGE
	INK_TC_YELLOW_W
	INK_TC_YELLOW_C
	INK_TC_LEMON
	INK_TC_GREEN_W
	INK_TC_GREEN_C
	INK_TC_TURQOUISE
	INK_TC_BLUE_LIGHT
	INK_TC_BLUE
	INK_TC_BLUE_DARK
	INK_TC_PURPLE
	INK_TC_VIOLET
	INK_TC_PINK

	_INK_MAX int = iota
)

// RGBA colors for each Ink
// Ordered exactly like the ink enum
var InkColorMap = [...]RGBA{
	{0, 0, 0, 0}, // INK_NONE

	{77, 56, 62, 255},    // INK_WRITE
	{102, 120, 142, 255}, // INK_CROSS
	{46, 71, 93, 255},    // INK_READ
	{146, 255, 99, 255},  // INK_BUFFER
	{255, 198, 99, 255},  // INK_AND
	{99, 242, 255, 255},  // INK_OR
	{174, 116, 255, 255}, // INK_XOR
	{255, 98, 138, 255},  // INK_NOT
	{255, 162, 0, 255},   // INK_NAND
	{48, 217, 255, 255},  // INK_NOR
	{166, 0, 255, 255},   // INK_XNOR
	{99, 255, 159, 255},  // INK_LATCH_ON
	{56, 77, 71, 255},    // INK_LATCH_OFF
	{255, 0, 65, 255},    // INK_CLOCK
	{255, 255, 255, 255}, // INK_LED
	{58, 69, 81, 255},    // INK_ANNOTATION
	{140, 171, 161, 255}, // INK_FILLER

	{42, 53, 65, 255},    // INK_TC_GRAY
	{159, 168, 174, 255}, // INK_TC_WHITE
	{161, 85, 94, 255},   // INK_TC_RED
	{161, 108, 86, 255},  // INK_TC_ORANGE
	{161, 133, 86, 255},  // INK_TC_YELLOW_W
	{161, 152, 86, 255},  // INK_TC_YELLOW_C
	{153, 161, 86, 255},  // INK_TC_LEMON
	{136, 161, 86, 255},  // INK_TC_GREEN_W
	{108, 161, 86, 255},  // INK_TC_GREEN_C
	{86, 161, 141, 255},  // INK_TC_TURQOUISE
	{86, 147, 161, 255},  // INK_TC_BLUE_LIGHT
	{86, 123, 161, 255},  // INK_TC_BLUE
	{86, 98, 161, 255},   // INK_TC_BLUE_DARK
	{102, 86, 161, 255},  // INK_TC_PURPLE
	{135, 86, 161, 255},  // INK_TC_VIOLET
	{161, 85, 151, 255},  // INK_TC_PINK
}

// Set Ink value for a given RGBA
func RGBAToInk(rgba RGBA) (Ink, error) {
	for i := 0; i < _INK_MAX; i++ {
		if rgba == InkColorMap[i] {
			return Ink(i), nil
		}
	}
	return Ink(0), fmt.Errorf("no ink found for rgba(%v)", rgba)
}

// Parses an image to an Ink Array
func ParseImageToInkArray(data []uint8, width int32, height int32) ([]Ink, error) {
	inks := make([]Ink, width*height)

	if len(data) != int(width*height*4) {
		return []Ink{}, fmt.Errorf("data size does not match given dimensions")
	}

	for i, j := 0, 0; i < len(data); i, j = i+4, j+1 {
		rgba := RGBA{data[i+0], data[i+1], data[i+2], data[i+3]}

		ink, err := RGBAToInk(rgba)
		if err != nil {
			return []Ink{}, fmt.Errorf("failed to convert color value (%v)", err)
		}

		inks[j] = ink
	}

	return inks, nil
}
