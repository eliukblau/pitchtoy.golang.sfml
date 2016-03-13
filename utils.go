package main

import (
	"math"
	"time"

	"fmt"

	sf "github.com/manyminds/gosfml"
)

var (
	deltaTime               = 0.0        // 1 segundo / 60 fps = 16ms (0.016ms) [tiempo por frame]
	lastFrameTime time.Time = time.Now() // entonces, timePerFrame = (time.Second / 60).Seconds()
)

var cursorRotation = 0.0

// actualiza el reloj interno para calculo de tiempo delta entre frames
func UpdateDeltaTime() {
	deltaTime = time.Since(lastFrameTime).Seconds()
	lastFrameTime = time.Now()
}

// dibuja los mensajes informativos en la escena principal
func DrawMessages(window *sf.RenderWindow, music *sf.Music, cursor *sf.Sprite) {
	if font, err := sf.NewFontFromFile(ResourcePath("font", "monoid.ttf")); err == nil {
		if text, err := sf.NewText(font); err == nil {
			text.SetCharacterSize(20)
			text.SetColor(sf.ColorWhite())
			text.SetPosition(sf.Vector2f{10, 10})

			angle := cursor.GetRotation()
			if angle > 180 {
				angle -= 360
			}
			text.SetString(fmt.Sprintf("Pitch: %.2f\n"+"Mouse X: %.f\n"+"Img. Angle: %.f\n", music.GetPitch(), cursor.GetPosition().X, angle))
			window.Draw(text, sf.DefaultRenderStates())
		}
	}
}

// realiza la animacion de balanceo del cursor segun el pitch de la musica
func AnimateCursor(cursor *sf.Sprite, pitch float32) {
	cursorRotation += deltaTime * math.Pow(float64(pitch+0.2), 5.0)
	angle := math.Cos(cursorRotation) * 50.0
	cursor.SetRotation(float32(angle))
}

// cambia Blue y Red del color de fondo segun la posicion del cursor
func UpdateBackColor(color *sf.Color, screenWidth uint, cursorPosX int) {
	dc := uint8(uint(cursorPosX) * 255 / screenWidth)
	color.B = 255 - dc
	color.R = dc
}

// interpolacion lineal del pitch de la musica segun la posicion del cursor
func UpdateMusicPitch(music *sf.Music, screenWidth uint, cursorPosX int) {
	music.SetPitch(float32(0.5 + math.Max(math.Min(float64(cursorPosX), float64(screenWidth)), 0.0)/float64(screenWidth)))
}
