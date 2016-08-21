package main

import (
	"runtime"

	"github.com/eliukblau/manami"
	sf "gitlab.com/tapir/sfml/v2.3/sfml"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	window := sf.NewRenderWindow(
		sf.VideoMode{Width: 1080, Height: 720, BitsPerPixel: 32},
		"PitchToy 2 (Go/SFML2)",
		sf.StyleClose|sf.StyleResize,
		&sf.ContextSettings{
			DepthBits:         0,
			StencilBits:       0,
			AntialiasingLevel: 0,
			MajorVersion:      2,
			MinorVersion:      0})

	manami.MainLoop(window, new(App), true)
}
