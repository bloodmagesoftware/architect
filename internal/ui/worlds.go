package ui

import (
	"github.com/bloodmagesoftware/architect/internal/app"
	mathf32 "github.com/bloodmagesoftware/architect/internal/math/f32"
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func worldsSidebarWidth() float32 {
	return mathf32.Max(10*vw, 8*rem)
}

func drawWorlds() {
	width := worldsSidebarWidth()
	height := 2 * rem
	rl.DrawTextEx(font, "Worlds", rl.Vector2{X: rem, Y: 0.5 * rem}, 16, 2, rl.Black)
	if rg.Button(rl.NewRectangle(0, height, width, height), "Create New") {
		if err := app.NewWorld(); err != nil {
			msgErr(err)
		}
	}
	for i, w := range app.Worlds {
		if rg.Button(rl.NewRectangle(0, float32(i+2)*height, width, height), w.Name) {
			if err := app.SetWorld(i); err != nil {
				msgErr(err)
			}
		}
	}
}
