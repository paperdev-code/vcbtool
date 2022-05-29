// Virtual Circuit Board Package
// Functions for parsing the image inside of blueprints
package vcb

// Ink is any of the possible usable logical inks
type Ink int

// Ink enum
// May change between version depending on game updates
const (
	InkNone Ink = iota
	InkTraceGrey
	InkTraceWhite
	InkTraceRed
	InkTraceOrange
	InkTraceYellowW
	InkTraceYellowC
	InkTraceLemon
	InkTraceGreenW
	InkTraceGreenC
	InkTraceTurqouise
	InkTraceLightBlue
	InkTraceBlue
	InkTraceDarkBlue
	InkTracePurple
	InkTraceViolet
	InkTracePink
	InkCross
	InkFiller
	InkAnnotation
	InkBuffer
	InkAND
	InkOR
	InkXOR
	InkNOT
	InkNAND
	InkNOR
	InkXNOR
	InkWrite
	InkRead
	InkLatchOn
	InkLatchOff
	InkClock
	InkLED
	_InkMax int = iota
)

// Convert Ink to Color
// Returns zero-ed Color and error on failure
func (ink Ink) ToColor() Color {
	if int(ink) >= _InkMax && int(ink) <= 0 {
		return Color{0, 0, 0, 0}
	}
	return Colors[int(ink)]
}

// No conversion, made for AnyPixel interface
func (ink Ink) ToInk() Ink {
	return ink
}
