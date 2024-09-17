package ui

import (
	"github.com/bloodmagesoftware/architect/internal/static"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	font  rl.Font
	vw    float32
	vwInt int32
	vh    float32
	vhInt int32
)

const (
	rem float32 = 16
)

func init() {
	rl.SetTraceLogLevel(rl.LogWarning)

	const hints uint32 = rl.FlagVsyncHint | rl.FlagMsaa4xHint | rl.FlagWindowResizable | rl.FlagWindowMaximized
	rl.SetConfigFlags(hints)
}

func MainLoop() {
	rl.InitWindow(800, 450, "architect")
	rl.MaximizeWindow()
	defer rl.CloseWindow()

	font = rl.LoadFontFromMemory(".ttf", static.NotoSans, 32, nil)
	defer rl.UnloadFont(font)

	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)
	rg.SetFont(font)

	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 16)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		layout()
		rl.ClearBackground(rl.RayWhite)

		drawWorlds()
		drawEditor()
		drawMsg()

		rl.EndDrawing()
	}
}

func layout() {
	vw = float32(rl.GetScreenWidth()) / 100
	vwInt = int32(rl.GetScreenWidth())
	vh = float32(rl.GetScreenHeight()) / 100
	vhInt = int32(rl.GetScreenHeight())
}
