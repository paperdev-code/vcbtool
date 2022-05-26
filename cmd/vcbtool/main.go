package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/charmbracelet/bubbletea"
	"github.com/paperdev-code/vcbtool/pkg/vcb"
)

func main() {
	bp, err := vcb.NewBlueprintFromBase64(os.Args[1])
	if err != nil {
		log.Fatalf("Error parsing blueprint (%v)", err)
	}

	cc, err := vcb.NewCircuitFromBlueprint(bp)
	if err != nil {
		log.Fatalf("Error parsing logic layer (%v)", err)
	}

	fmt.Print("Circuit:")
	for i := 0; i < len(cc.Logic_layer); i += 1 {
		if i%int(cc.Width) == 0 {
			fmt.Print("\n")
		}
		rgba := vcb.InkColors[cc.Logic_layer[i]]
		color_bg := fmt.Sprintf("\x1b[48;2;%d;%d;%dm", rgba.R, rgba.G, rgba.B)
		color_fg := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", rgba.R/2, rgba.G/2, rgba.B/2)
		fmt.Printf("%s%s%.2d\x1b[0m", color_bg, color_fg, cc.Logic_layer[i])
	}
	fmt.Print("\n")
}
