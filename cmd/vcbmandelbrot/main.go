package main

import (
	"fmt"
	"log"
	"math"

	"github.com/paperdev-code/vcbtool/pkg/vcb"
)

const Width, Height = 120, 120
const MaxIterations = 40

func main() {
	layer, err := vcb.NewLayer[vcb.LogicLayer](Width, Height)
	if err != nil {
		log.Fatalf("Failed to create layer; %v", layer)
	}

	ctx := vcb.NewContext(layer)

	for y := 0; y < layer.Height; y++ {
		for x := 0; x < layer.Width; x++ {

			// Map to range -2 to 2
			a := float64(x)/float64(layer.Width)*3.5 - 2.5
			b := float64(y)/float64(layer.Height)*4.5 - 2.25
			ca := a
			cb := b

			iterations := 0
			for iterations < MaxIterations {
				aa := a*a - b*b
				bb := 2 * a * b
				a = aa + ca
				b = bb + cb

				if math.Abs(a+b) > 16 {
					break
				}

				iterations += 1
			}

			ink := vcb.InkNone
			if iterations != MaxIterations {
				i := float64(iterations) / float64(MaxIterations) * float64(vcb.InkTracePink-vcb.InkTraceGrey)
				ink = vcb.InkTraceGrey + vcb.Ink(i) - 1
			}
			ctx.Set(ink, x, y)
		}
	}

	bp, err := vcb.NewBlueprintFromLayer(layer)
	if err != nil {
		log.Fatalf("Failed to convert layer to blueprint; %v", err)
	}

	b64, err := bp.ToBase64()
	if err != nil {
		log.Fatalf("Failed to convert blueprint to base64; %v", err)
	}
	fmt.Printf("%s\n", b64)
}
