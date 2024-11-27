package ui

import (
	"github.com/bloodmagesoftware/architect/internal/app"
	mathf32 "github.com/bloodmagesoftware/architect/internal/math/f32"
	"github.com/bloodmagesoftware/architect/internal/ui/ctp"
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func worldsSidebarWidth() float32 {
	return mathf32.Max(10*vw, 8*rem)
}

func drawWorldsList() {
	height := vh - atlasHeight()
	width := worldsSidebarWidth()

	entryHeight := 2 * rem

	rl.DrawRectangleLines(0, 0, int32(width), int32(height), ctp.Mantle)
	rl.DrawTextEx(font, "Worlds", rl.Vector2{X: rem, Y: 0.5 * rem}, 16, 2, ctp.Text)
	if rg.Button(rl.NewRectangle(0, entryHeight, width, entryHeight), "Create New") {
		if err := app.NewWorld(); err != nil {
			msgErr(err)
		}
	}

	for i, w := range app.Worlds {
		startY := float32(i+2) * entryHeight
		if rg.Button(rl.NewRectangle(0, startY, width, entryHeight), w.Name) {
			if err := app.SetWorld(i); err != nil {
				msgErr(err)
			}
		}
	}
}
