package ui

import (
	"github.com/bloodmagesoftware/architect/internal/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawEditor() {
	if app.World < 0 {
		return
	}

	wsw := int32(worldsSidebarWidth())
	rl.DrawRectangle(wsw, 0, vwInt-wsw, vhInt, rl.Black)
}
