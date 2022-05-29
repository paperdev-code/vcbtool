package vcb

import (
	"testing"
)

const TEST_BLUEPRINT = "KLUv/SCQnQMAEkgaI5ArAMhhSWZsnfijT+C5S5JrveAkYl1y7xYWkeObps5JUmkfdzmRDwbC0b00QWO3P817BF/mIdWvyP7Tbp1vU0ZIgScL6wI0baF66bi0bsXL4K1wSTl7zT5mMfPS2ei8eoLq2QY7GgACAgDADg+AGQIAAAAGAAAAAgAAAAYAAAACAAAAkAAAAAIAAAAAAAEA"

var TEST_LOGIC_LAYER = []Ink{
	InkTraceGrey, InkTraceWhite, InkTraceRed, InkTraceOrange, InkTraceYellowW, InkTraceYellowC,
	InkTraceLemon, InkTraceGreenW, InkTraceGreenC, InkTraceTurqouise, InkTraceLightBlue, InkTraceBlue,
	InkTraceDarkBlue, InkTracePurple, InkTraceViolet, InkTracePink, InkWrite, InkCross,
	InkRead, InkBuffer, InkAND, InkOR, InkXOR, InkNOT,
	InkNAND, InkNOR, InkXNOR, InkLatchOn, InkLatchOff, InkClock,
	InkLED, InkAnnotation, InkFiller, InkNone, InkNone, InkNone,
}

var TEST2_LOGIC_LAYER = []Ink{
	InkTraceGrey, InkTraceWhite, InkTraceRed, InkTraceOrange, InkTraceYellowW, InkTraceYellowC,
	InkTraceLemon, InkTraceGreenW, InkTraceGreenC, InkTraceTurqouise, InkTraceLightBlue, InkTraceBlue,
	InkTraceDarkBlue, InkTracePurple, InkTraceViolet, InkTracePink, InkWrite, InkCross,
	InkRead, InkBuffer, InkAND, InkOR, InkXOR, InkNOT,
	InkNAND, InkNOR, InkXNOR, InkLatchOn, InkLatchOff, InkClock,
	InkLED, InkAnnotation, InkFiller, InkLED, InkLED, InkLED,
}

// This test decodes from base64 from the game
// Checks whether the circuit was parsed correctly
// Draws a line of LEDs on it
// Then parses the circuit back to base64
// Checks whether the circuit was parsed correctly
func TestCircuitFromBase64(t *testing.T) {
	bp, err := NewBlueprintFromBase64(TEST_BLUEPRINT)
	if err != nil {
		t.Fatalf("Fail: %v", err)
	}

	circuit := bp.Layer.(*LogicLayer).Circuit
	for i, ink := range circuit {
		if ink != TEST_LOGIC_LAYER[i] {
			t.Fatalf("Wrong color on index %d, should be %d", i, TEST_LOGIC_LAYER[i])
		}
	}

	ctx := NewContext(bp.Layer)
	ctx.Line(InkLED, 4, 6, 6, 6)

	b64, err := bp.ToBase64()
	if err != nil {
		t.Fatalf("Fail: %v", err)
	}

	cb, err := NewBlueprintFromBase64(b64)
	if err != nil {
		t.Fatalf("Fail: %v", err)
	}

	circuit = cb.Layer.(*LogicLayer).Circuit
	for i, ink := range circuit {
		if ink != TEST_LOGIC_LAYER[i] {
			t.Fatalf("Wrong color on index %d, should be %d", i, TEST2_LOGIC_LAYER[i])
		}
	}
}
