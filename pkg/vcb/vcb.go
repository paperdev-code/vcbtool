// Virtual Circuit Board package
// Data structures and methods to alter or create VCB blueprints
package vcb

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/DataDog/zstd"
)

// Blueprints are images encoded as base64
// This is the data structure encoded within
// (Shoutout chrjen's vcbb blueprint impl)
type Blueprint struct {
	Data   []uint8
	Footer struct {
		HeightType int32
		Height     int32
		WidthType  int32
		Width      int32
		BytesType  int32
		Bytes      int32
		LayerType  int32
		Layer      int32
	}
}

// A data structure containing the layers
type Circuit struct {
	Width       int32
	Height      int32
	Logic_layer []Ink
	//TODO add support for decorative layers
	//deco_layer []uint8
}

// Creates a blueprint from a base64 encoded string
func NewBlueprintFromBase64(b64 string) (Blueprint, error) {
	blueprint := Blueprint{}

	input := strings.NewReader(b64)
	decoded, err := io.ReadAll(base64.NewDecoder(base64.StdEncoding, input))
	if err != nil {
		return Blueprint{}, fmt.Errorf("failed to decode and read base64 blueprint. (%v)", err)
	}

	footer_size := binary.Size(blueprint.Footer)
	binary.Read(bytes.NewReader(decoded[len(decoded)-footer_size:]), binary.LittleEndian, &blueprint.Footer)

	decompressed, err := zstd.Decompress(nil, decoded[:len(decoded)-footer_size])
	if err != nil {
		return Blueprint{}, fmt.Errorf("failed to decompress blueprint layer. (%v)", err)
	}
	blueprint.Data = decompressed

	return blueprint, nil
}

// Creates a circuit from a blueprint
func NewCircuitFromBlueprint(blueprint Blueprint) (Circuit, error) {
	circuit := Circuit{
		Width:  blueprint.Footer.Width,
		Height: blueprint.Footer.Height,
	}

	logic_layer, err := ParseImageToInkArray(
		blueprint.Data,
		circuit.Width,
		circuit.Height,
	)
	if err != nil {
		return Circuit{}, fmt.Errorf("failed to parse logic layer (%v)", err)
	}
	circuit.Logic_layer = logic_layer

	return circuit, nil
}

func NewBlueprintFromCircuit(circuit Circuit) (Blueprint, error) {
	blueprint := Blueprint{}

	return blueprint, nil
}
