package ui

import (
	"github.com/bloodmagesoftware/architect/internal/file"
	mathf32 "github.com/bloodmagesoftware/architect/internal/math/f32"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func atlasHeight() float32 {
	return mathf32.Max(10*vh, 8*rem)
}

func drawAtlas() {
	width := vwInt
	height := int32(atlasHeight())
	yStart := vhInt - height

	rl.DrawRectangleLines(0, yStart, width, height, rl.Black)

	entryWidth := width / 8
	entryHeight := height / 2

	for i, ae := range file.AtlasEntries {
		x := int32(i % 8)
		y := int32(i / 8)
		ae.Draw(font, x*entryWidth, yStart+y*entryHeight, entryWidth, entryHeight)
	}
}
