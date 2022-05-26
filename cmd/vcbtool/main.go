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
		if i%int(cc.Height) == 0 {
			fmt.Print("\n")
		}
		fmt.Printf("[%.2d]", cc.Logic_layer[i])
	}
	fmt.Print("\n")
}
