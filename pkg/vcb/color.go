// Colors are RGBA values for Decor Layers and Ink conversion
package vcb

import "fmt"

// Colors contain RGBA values
type Color struct {
	R, G, B, A byte
}

// Image type
type Image []Color

// Convert Color to Ink
// Returns INK_NONE and Error on failure
func (c Color) ToInk() Ink {
	for i, color := range Colors {
		if c == color {
			return Ink(i)
		}
	}
	return InkNone
}

// No conversion, made for AnyPixel interface
func (c Color) ToColor() Color {
	return c
}

// Convert raw bytes to Color
func ColorFromData(data []byte) (Color, error) {
	if len(data) != 4 {
		return Color{0, 0, 0, 0},
			fmt.Errorf("data must have len() = 4")
	}
	return Color{
		R: data[0],
		G: data[1],
		B: data[2],
		A: data[3],
	}, nil
}

func (image Image) ToData() []byte {
	data := make([]uint8, len(image)*4)
	for i, j := 0, 0; i < len(image); i, j = i+1, j+4 {
		data[j+0], data[j+1], data[j+2], data[j+3] = image[i].R, image[i].G, image[i].B, image[i].A
	}
	return data
}

// RGBA colors for each Ink
// Ordered exactly like the ink enum
var Colors = [...]Color{
	{0, 0, 0, 0},         // InkNone
	{42, 53, 65, 255},    // InkTraceGrey
	{159, 168, 174, 255}, // InkTraceWhite
	{161, 85, 94, 255},   // InkTraceRed
	{161, 108, 86, 255},  // InkTraceOrange
	{161, 133, 86, 255},  // InkTraceYellowW
	{161, 152, 86, 255},  // InkTraceYellowC
	{153, 161, 86, 255},  // InkTraceLemon
	{136, 161, 86, 255},  // InkTraceGreenW
	{108, 161, 86, 255},  // InkTraceGreenC
	{86, 161, 141, 255},  // InkTraceTurqouise
	{86, 147, 161, 255},  // InkTraceLightBlue
	{86, 123, 161, 255},  // InkTraceBlue
	{86, 98, 161, 255},   // InkTraceDarkBlue
	{102, 86, 161, 255},  // InkTracePurple
	{135, 86, 161, 255},  // InkTraceViolet
	{161, 85, 151, 255},  // InkTracePink
	{102, 120, 142, 255}, // InkCross
	{140, 171, 161, 255}, // InkFiller
	{58, 69, 81, 255},    // InkAnnotation
	{146, 255, 99, 255},  // InkBuffer
	{255, 198, 99, 255},  // InkAND
	{99, 242, 255, 255},  // InkOR
	{174, 116, 255, 255}, // InkXOR
	{255, 98, 138, 255},  // InkNOT
	{255, 162, 0, 255},   // InkNAND
	{48, 217, 255, 255},  // InkNOR
	{166, 0, 255, 255},   // InkXNOR
	{77, 56, 62, 255},    // InkWrite
	{46, 71, 93, 255},    // InkRead
	{99, 255, 159, 255},  // InkLatchOn
	{56, 77, 71, 255},    // InkLatchOff
	{255, 0, 65, 255},    // InkClock
	{255, 255, 255, 255}, // InkLED
}
