package ui

import (
	"github.com/bloodmagesoftware/architect/internal/file"
	"github.com/bloodmagesoftware/architect/internal/static"
	"github.com/bloodmagesoftware/architect/internal/ui/ctp"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	font         rl.Font
	vw           float32
	vwInt        int32
	vh           float32
	vhInt        int32
	blockSize    float32
	blockSizeInt int32
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

	rg.SetStyle(rg.DEFAULT, rg.BASE_COLOR_NORMAL, 0x313244ff)
	rg.SetStyle(rg.DEFAULT, rg.TEXT_COLOR_NORMAL, 0xcdd6f4ff)
	rg.SetStyle(rg.DEFAULT, rg.BORDER_COLOR_NORMAL, 0x1e1e2eff)
	rg.SetStyle(rg.DEFAULT, rg.BASE_COLOR_FOCUSED, 0x6c7086ff)
	rg.SetStyle(rg.DEFAULT, rg.TEXT_COLOR_FOCUSED, 0xcdd6f4ff)
	rg.SetStyle(rg.DEFAULT, rg.BORDER_COLOR_FOCUSED, 0x1e1e2eff)
	rg.SetStyle(rg.DEFAULT, rg.BASE_COLOR_PRESSED, 0x7f849cff)
	rg.SetStyle(rg.DEFAULT, rg.TEXT_COLOR_PRESSED, 0xcdd6f4ff)
	rg.SetStyle(rg.DEFAULT, rg.BORDER_COLOR_PRESSED, 0x1e1e2eff)

	rl.SetTargetFPS(30)

	if err := file.UpdateAtlas(); err != nil {
		msgErr(err)
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		layout()
		rl.ClearBackground(ctp.Base)

		drawWorldsList()
		drawAtlas()
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
	blockSize = 4 * vw
	blockSizeInt = int32(blockSize)
}
