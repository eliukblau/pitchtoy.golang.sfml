package main

import (
	sf "gitlab.com/tapir/sfml/v2.3/sfml"
)

type App struct {
	isRunning bool
	window    *sf.RenderWindow
	backColor sf.Color
}

func (app *App) Init(window *sf.RenderWindow) {
	app.window = window
	app.isRunning = true

	app.backColor = sf.Color{R: 127, G: 32, B: 127, A: 255}

	// activamos la sincronizacion vertical
	app.window.SetVerticalSyncEnabled(true)

	// quitamos el cursor del sistema operativo
	app.window.SetMouseCursorVisible(false)
}

func (app *App) Input(event *sf.Event) {
	switch event.Type {
	case sf.EventClosed:
		app.isRunning = false

	case sf.EventKeyReleased:
		switch event.Key.Code {
		case sf.KeyEscape:
			app.isRunning = false

		}
	}
}

func (app *App) Update(dt float32) {}

func (app *App) Draw(alpha float32) {
	app.window.Clear(app.backColor)
}

func (app *App) Quit() {}

func (app *App) IsRunning() bool {
	return app.isRunning
}
