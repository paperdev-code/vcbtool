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

	"github.com/paperdev-code/vcbtool/pkg/zstd"
)

// Blueprints contain a single layer
// Depending on which Type is set
type Blueprint struct {
	Width, Height int
	Layer         AnyLayer
	Type          LayerType
}

// Layer type
type LayerType int

// Layer types
const (
	LayerTypeLogic    LayerType = 65536
	LayerTypeDecorOn  LayerType = 131072
	LayerTypeDecorOff LayerType = 262144
)

// Binary representation of the footer
type BlueprintFooter struct {
	HeightPadding int32
	Height        int32
	WidthPadding  int32
	Width         int32
	BytesPadding  int32
	Bytes         int32
	LayerPadding  int32
	Layer         int32
}

// Create a new blueprint from layer
func NewBlueprintFromLayer(layer AnyLayer) (*Blueprint, error) {
	bp := Blueprint{}
	err := bp.SetLayer(layer)
	if err != nil {
		return nil, fmt.Errorf("failed to create blueprint, %v", err)
	}
	return &bp, err
}

// Set the blueprint layer
func (bp *Blueprint) SetLayer(layer AnyLayer) error {
	switch (layer).(type) {
	case *LogicLayer:
		bp.Type = LayerTypeLogic
	case *DecorOnLayer:
		bp.Type = LayerTypeDecorOn
	case *DecorOffLayer:
		bp.Type = LayerTypeDecorOff
	default:
		return fmt.Errorf("'%T' is not a valid layer type", layer)
	}

	bp.Width, bp.Height = AnyLayer.GetDimensions(layer)
	bp.Layer = layer
	return nil
}

// Convert into a base64 representation
// Allows to export blueprints to the game
func (bp *Blueprint) ToBase64() (string, error) {
	image, err := bp.Layer.ToImage()
	if err != nil {
		return "", fmt.Errorf("failed to convert image; %v", err)
	}

	var data []uint8

	// Compress image data
	{
		data, err = zstd.Compress(image.ToData())
		if err != nil {
			return "", fmt.Errorf("failed to compress image; %v", err)
		}
	}

	footer := BlueprintFooter{
		HeightPadding: 2,
		Height:        int32(bp.Height),
		WidthPadding:  2,
		Width:         int32(bp.Width),
		BytesPadding:  2,
		Bytes:         int32(bp.Width * bp.Height * 4),
		LayerPadding:  2,
		Layer:         int32(bp.Type),
	}

	var output strings.Builder

	// Encode to base64
	{
		encoder := base64.NewEncoder(base64.StdEncoding, &output)

		_, err = encoder.Write(data)
		if err != nil {
			return "", fmt.Errorf("failed to encode data; %v", err)
		}

		err = binary.Write(encoder, binary.LittleEndian, footer)
		if err != nil {
			return "", fmt.Errorf("failed to encode footer; %v", err)
		}
		encoder.Close()
	}

	return output.String(), nil
}

// Create a new blueprint from a base64 encoded string
// Allows to import blueprints from the game
func NewBlueprintFromBase64(b64 string) (*Blueprint, error) {
	bp := Blueprint{}

	var data []byte = nil
	var footer_size int = 0

	// Decode base64 into binary
	{
		reader := strings.NewReader(b64)
		var err error = nil
		data, err = io.ReadAll(base64.NewDecoder(base64.StdEncoding, reader))
		if err != nil {
			return nil, fmt.Errorf("failed to decode base64; %v", err)
		}
	}

	var layer_type LayerType

	// Extract footer from base64 encoded string
	{
		footer := BlueprintFooter{}
		footer_size = binary.Size(footer)

		binary.Read(
			bytes.NewReader(data[len(data)-footer_size:]),
			binary.LittleEndian,
			&footer,
		)

		bp.Width = int(footer.Width)
		bp.Height = int(footer.Height)
		bp.Type = LayerType(footer.Layer)
	}

	colors := make([]Color, bp.Width*bp.Height)

	// Extract image
	{
		image, err := zstd.Decompress(data[:len(data)-footer_size])
		if err != nil {
			return nil, fmt.Errorf("failed to decompress data; %v", err)
		}

		if len(image) != len(colors)*4 {
			return nil, fmt.Errorf("image and footer size mismatch")
		}

		for i, j := 0, 0; i < len(colors); i, j = i+1, j+4 {
			// Last check ensures data is divisible by 4, no need for error check
			colors[i], _ = ColorFromData(image[j : j+4])
		}
	}

	var err error
	switch bp.Type {
	case LayerTypeLogic:
		bp.Layer, _ = NewLayer[LogicLayer](bp.Width, bp.Height)
		err = bp.Layer.FromImage(colors)
	case LayerTypeDecorOn:
		bp.Layer, _ = NewLayer[DecorOnLayer](bp.Width, bp.Height)
		err = bp.Layer.FromImage(colors)
	case LayerTypeDecorOff:
		bp.Layer, _ = NewLayer[DecorOffLayer](bp.Width, bp.Height)
		err = bp.Layer.FromImage(colors)
	default:
		return nil, fmt.Errorf("invalid layer type (%d)", layer_type)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to set blueprint data; %v", err)
	}

	return &bp, nil
}
