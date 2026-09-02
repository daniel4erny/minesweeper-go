package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"os"
)

func main(){
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Println("idk, je to v pici")
		os.Exit(1)
	}

	err = screen.Init()
	if err != nil {
		fmt.Println("idk, je to v pici")
		os.Exit(1)
	}

	defer screen.Fini()

	screen.SetStyle(tcell.StyleDefault)

	for {
		screen.PutStr(1,1, "BALLS")
		screen.Show()
	}
}