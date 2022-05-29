package main

import (
	"fmt"
	"log"
	"os"

	"github.com/paperdev-code/vcbtool/pkg/vcb"
)

func main() {
	bp, err := vcb.NewBlueprintFromBase64(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to decode blueprint; %v.", err)
	}

	cc, ok := bp.Layer.(*vcb.LogicLayer)
	if !ok {
		log.Fatalf("Expected blueprint for Logic Layer.\n")
	}

	fmt.Print("Circuit:")
	for i, ink := range cc.Circuit {
		if i%int(cc.Width) == 0 {
			fmt.Print("\n")
		}
		if ink != vcb.InkNone {
			color := vcb.Colors[ink]
			color_bg := fmt.Sprintf("\x1b[48;2;%d;%d;%dm", color.R, color.G, color.B)
			color_fg := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", color.R/2, color.G/2, color.B/2)
			fmt.Printf("%s%s%.2d\x1b[0m", color_bg, color_fg, ink)
		} else {
			fmt.Printf("  ")
		}
	}
	fmt.Print("\n")

	b64, err := bp.ToBase64()
	if err != nil {
		log.Fatalf("Failed to convert to base64; %v.", err)
	}
	fmt.Printf("Blueprint:\n%s\n", b64)
}
