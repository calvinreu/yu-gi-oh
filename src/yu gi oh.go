package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	var config Config

	args := os.Args[1:]
	if len(args) > 2 {
		fmt.Println("to many aunch arguments")
	}
	if len(args) < 1 {
		fmt.Println("to few arguments")
		return
	}

	err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_EVENTS)
	if err != nil {
		fmt.Println(err)
		return
	}

	config.Load(args[0])
	gameloop(config.Convert())

	sdl.Quit()
}
