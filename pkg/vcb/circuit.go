package vcb

import "fmt"

// A data structure containing the layers
type Circuit struct {
	Width  int32
	Height int32
	Layer  []Ink
}

// Creates a circuit from a blueprint
func NewCircuitFromBlueprint(blueprint Blueprint) (Circuit, error) {
	circuit := Circuit{
		Width:  blueprint.Footer.Width,
		Height: blueprint.Footer.Height,
	}

	layer, err := ParseRGBAToInk(
		blueprint.Data,
		circuit.Width,
		circuit.Height,
	)
	if err != nil {
		return Circuit{}, fmt.Errorf("failed to parse logic layer; %v", err)
	}
	circuit.Layer = layer

	return circuit, nil
}
