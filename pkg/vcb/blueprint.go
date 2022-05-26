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

	"github.com/klauspost/compress/zstd"
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

// Creates a blueprint from a base64 encoded string
func NewBlueprintFromBase64(b64 string) (Blueprint, error) {
	blueprint := Blueprint{}

	input := strings.NewReader(b64)
	decoded, err := io.ReadAll(base64.NewDecoder(base64.StdEncoding, input))
	if err != nil {
		return Blueprint{}, fmt.Errorf("failed to decode and read base64 blueprint; %v", err)
	}

	footer_size := binary.Size(blueprint.Footer)
	binary.Read(bytes.NewReader(decoded[len(decoded)-footer_size:]), binary.LittleEndian, &blueprint.Footer)

	if blueprint.Footer.Layer != 65536 {
		return Blueprint{}, fmt.Errorf("blueprint should be logic '65536' type, is '%d'", blueprint.Footer.Layer)
	}

	zstd_reader, err := zstd.NewReader(nil)
	if err != nil {
		return Blueprint{}, fmt.Errorf("reader fail; %v", err)
	}

	decompressed, err := zstd_reader.DecodeAll(decoded[:len(decoded)-footer_size], nil)
	if err != nil {
		return Blueprint{}, fmt.Errorf("failed to decompress blueprint layer; %v", err)
	}
	blueprint.Data = decompressed

	return blueprint, nil
}

// Reverse operation that builds an image from a circuit
func NewBlueprintFromCircuit(circuit Circuit) (Blueprint, error) {
	blueprint := Blueprint{}

	blueprint.Footer.HeightType = 2
	blueprint.Footer.WidthType = 2
	blueprint.Footer.LayerType = 2

	blueprint.Footer.Width = circuit.Width
	blueprint.Footer.Height = circuit.Height
	blueprint.Footer.Layer = 0
	data, err := ParseInkToRGBA(circuit.Layer, circuit.Width, circuit.Height)
	if err != nil {
		return Blueprint{}, fmt.Errorf("failed to create blueprint from circuit; %v", err)
	}

	blueprint.Data = data
	return blueprint, nil
}

func Base64FromBlueprint(blueprint Blueprint) (string, error) {
	b64 := ""
	return b64, nil
}
