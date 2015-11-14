package main

import (
	"runtime"
	"time"

	sf "github.com/manyminds/gosfml"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	// 0 - CONSTANTES Y VARIABLES GLOBALES

	const (
		winWidth  = 1080
		winHeight = 720
	)

	var (
		backColor = sf.Color{R: 127, G: 32, B: 127, A: 255}
		pitch     = float32(1.0)
	)

	// 1 - INICIAR EL JUEGO

	// creamos la ventana del juego
	window := sf.NewRenderWindow(
		sf.VideoMode{Width: winWidth, Height: winHeight, BitsPerPixel: 32},
		"Pitch Toy (Go/SFML2)",
		sf.StyleTitlebar|sf.StyleClose,
		sf.DefaultContextSettings())

	// cerrado de ventana al final de la ejecucion
	defer window.Close()

	// activamos la sincronizacion vertical
	window.SetVSyncEnabled(true)

	// quitamos el cursor del sistema operativo
	window.SetMouseCursorVisible(false)

	// centramos la ventana
	window.SetPosition(sf.Vector2i{
		X: int((sf.GetDesktopVideoMode().Width - window.GetSize().X) / 2),
		Y: int((sf.GetDesktopVideoMode().Height - window.GetSize().Y) / 2),
	})

	// cargamos la textura de la imagen de fondo
	texture, err := sf.NewTextureFromFile(ResourcePath("gfx", "cursor.png"), nil)
	if err != nil {
		panic(err)
	}
	// creamos un sprite a partir de la textura
	cursor, err := sf.NewSprite(texture)
	if err != nil {
		panic(err)
	}
	// hacemos que la textura se vea suave al transformarla
	cursor.GetTexture().SetSmooth(true)
	// establecemos el origen y las coordenadas iniciales del sprite
	cursor.SetOrigin(sf.Vector2f{X: 150, Y: 150})
	cursor.SetPosition(sf.Vector2f{
		X: float32(window.GetSize().X) / 2.0,
		Y: float32(window.GetSize().Y) / 2.0,
	})

	// cargamos la musica de fondo
	music, err := sf.NewMusicFromFile(ResourcePath("sfx", "music.ogg"))
	if err != nil {
		panic(err)
	}
	// reproducimos la musica de fondo
	music.SetLoop(true)
	music.Play()

	// 2 - BUCLE PRINCIPAL DEL JUEGO

	gameloop := true
	for gameloop && window.IsOpen() {
		// 2.1 - PROCESA LA ENTRADA Y...
		// 2.2 - ACTUALIZAR EL ESTADO DEL JUEGO
		for event := window.PollEvent(); event != nil; event = window.PollEvent() {
			switch ev := event.(type) {
			case sf.EventClosed:
				gameloop = false

			case sf.EventKeyReleased:
				switch ev.Code {
				case sf.KeyEscape:
					gameloop = false
				}

			case sf.EventMouseMoved:
				// color de fondo
				colorDelta := uint8(uint(ev.X) * 255 / window.GetSize().X)
				backColor.R = colorDelta
				backColor.B = 255 - colorDelta

				// posicion del cursor
				cursor.SetPosition(sf.Vector2f{
					X: float32(ev.X),
					Y: float32(ev.Y),
				})

				// pitch de la musica de fondo
				pitch = float32(ev.X) / (float32(window.GetSize().X) / 2.0)
				if pitch <= 0.0 {
					pitch = 0.01
				} else if pitch >= 1.998 {
					pitch = 2.0
				}
				music.SetPitch(float32(pitch))
			}
		}

		// animamos la rotacion del cursor
		dt := (time.Second / 60).Seconds()
		cursor.Rotate(180 * pitch * float32(dt))

		// 2.3 - RENDERIZAR EL JUEGO
		window.Clear(backColor)
		window.Draw(cursor, sf.DefaultRenderStates())
		window.Display()
	}

	// 3 - FINALIZAR EL JUEGO
	// defer's de go ya se encargan de liberar y finalizar todo! ;)

}
