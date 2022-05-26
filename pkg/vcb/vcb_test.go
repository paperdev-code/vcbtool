package vcb

import (
	"testing"
)

const TEST_BLUEPRINT = "KLUv/SCQnQMAEkgaI5ArAMhhSWZsnfijT+C5S5JrveAkYl1y7xYWkeObps5JUmkfdzmRDwbC0b00QWO3P817BF/mIdWvyP7Tbp1vU0ZIgScL6wI0baF66bi0bsXL4K1wSTl7zT5mMfPS2ei8eoLq2QY7GgACAgDADg+AGQIAAAAGAAAAAgAAAAYAAAACAAAAkAAAAAIAAAAAAAEA"

var TEST_LOGIC_LAYER = []Ink{
	INK_TC_GRAY, INK_TC_WHITE, INK_TC_RED, INK_TC_ORANGE, INK_TC_YELLOW_W, INK_TC_YELLOW_C,
	INK_TC_LEMON, INK_TC_GREEN_W, INK_TC_GREEN_C, INK_TC_TURQOUISE, INK_TC_BLUE_LIGHT, INK_TC_BLUE,
	INK_TC_BLUE_DARK, INK_TC_PURPLE, INK_TC_VIOLET, INK_TC_PINK, INK_WRITE, INK_CROSS,
	INK_READ, INK_BUFFER, INK_AND, INK_OR, INK_XOR, INK_NOT,
	INK_NAND, INK_NOR, INK_XNOR, INK_LATCH_ON, INK_LATCH_OFF, INK_CLOCK,
	INK_LED, INK_ANNOTATION, INK_FILLER, INK_NONE, INK_NONE, INK_NONE,
}

func TestCircuitFromBase64(t *testing.T) {
	bp, err := NewBlueprintFromBase64(TEST_BLUEPRINT)
	if err != nil {
		t.Fatalf("NewBlueprintFromBase64() failure; %v", err)
	}

	// Check whether footer parsed correctly.
	{
		valid := true

		if bp.Footer.Width != 6 {
			valid = false
		}

		if bp.Footer.Height != 6 {
			valid = false
		}

		if !valid {
			t.Fatalf(
				"Blueprint dimensions are (%d, %d), should be (6, 6)",
				bp.Footer.Width,
				bp.Footer.Height,
			)
		}
	}

	cc, err := NewCircuitFromBlueprint(bp)
	if err != nil {
		t.Fatalf("NewCircuitFromBlueprint() failure; %v", err)
	}

	// Check whether all ink was converted properly
	{
		x := int(0)
		y := int(0)
		for i := 0; i < len(cc.Layer); i += 1 {
			correct_ink := TEST_LOGIC_LAYER[i]
			circuit_ink := cc.Layer[i]
			if i%int(cc.Width) == 0 && i != 0 {
				y += 1
			}
			x += 1
			if x == int(cc.Width) {
				x = 0
			}
			if cc.Layer[i] != TEST_LOGIC_LAYER[i] {
				t.Fatalf("Ink @ (%d, %d) is %d, should be %d",
					x,
					y,
					circuit_ink,
					correct_ink,
				)
			}
		}
	}

	// Reconstruct blueprint from circuit
	rbp, err := NewBlueprintFromCircuit(cc)
	if err != nil {
		t.Fatalf("NewBlueprintFromCircuit() failure; %v", err)
	}

	//
	{
		_ = rbp
	}
}
