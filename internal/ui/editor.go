package ui

import (
	"github.com/bloodmagesoftware/architect/internal/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawEditor() {
	if app.World < 0 {
		return
	}

	startX := int32(worldsSidebarWidth())
	width := vwInt - startX
	height := vhInt - int32(atlasHeight())

	rl.DrawRectangle(startX, 0, width, height, rl.Black)

	x := rl.GetMouseX()
	y := rl.GetMouseY()
	if x >= startX && y >= 0 && x < vwInt && y < height {
		xRel := (x - startX) / blockSizeInt
		yInt := y / blockSizeInt
		rl.DrawRectangle(xRel*blockSizeInt+startX, yInt*blockSizeInt, blockSizeInt, blockSizeInt, rl.Fade(rl.Gray, 0.5))
	}
}
