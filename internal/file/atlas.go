package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/bloodmagesoftware/architect/internal/ui/ctp"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type (
	AtlasEntry interface {
		Draw(font rl.Font, posX, posY, width, height int32)
	}

	AtlasObject struct {
		Texture string
	}

	AtlasDirectory struct {
		relativePath string
	}
)

var (
	AtlasPath    = "./assets"
	AtlasEntries []AtlasEntry
)

func UpdateAtlas() error {
	de, err := os.ReadDir(AtlasPath)
	if err != nil {
		return errors.Join(errors.New("failed to read directory "+AtlasPath), err)
	}

	ae := []AtlasEntry{AtlasDirectory{relativePath: ".."}}
	for _, entry := range de {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		if entry.IsDir() {
			ae = append(ae, AtlasDirectory{relativePath: entry.Name()})
		} else {
			ae = append(ae, AtlasObject{Texture: filepath.Join(AtlasPath, entry.Name())})
		}
	}
	AtlasEntries = ae
	return nil
}

func (d AtlasDirectory) Draw(font rl.Font, posX, posY, width, height int32) {
	mouseX := rl.GetMouseX()
	mouseY := rl.GetMouseY()

	if rl.CheckCollisionPointRec(rl.Vector2{X: float32(mouseX), Y: float32(mouseY)}, rl.NewRectangle(float32(posX), float32(posY), float32(width), float32(height))) {
		rl.DrawRectangle(posX, posY, width, height, rl.Fade(ctp.Surface1, 0.5))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			rl.DrawRectangle(posX, posY, width, height, ctp.Surface1)
			AtlasPath = filepath.Join(AtlasPath, d.relativePath)
			_ = UpdateAtlas()
		} else {
			rl.DrawRectangle(posX, posY, width, height, ctp.Surface0)
		}
	}
	rl.DrawRectangleLines(posX, posY, width, height, ctp.Mantle)
	rl.DrawTextEx(font, d.relativePath, rl.Vector2{X: float32(posX), Y: float32(posY)}, 16, 2, ctp.Text)
}

func (d AtlasObject) Draw(font rl.Font, posX, posY, width, height int32) {
	mouseX := rl.GetMouseX()
	mouseY := rl.GetMouseY()

	if rl.CheckCollisionPointRec(rl.Vector2{X: float32(mouseX), Y: float32(mouseY)}, rl.NewRectangle(float32(posX), float32(posY), float32(width), float32(height))) {
		rl.DrawRectangle(posX, posY, width, height, rl.Fade(ctp.Surface1, 0.5))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			rl.DrawRectangle(posX, posY, width, height, ctp.Surface1)
		} else {
			rl.DrawRectangle(posX, posY, width, height, ctp.Surface0)
		}
	}

	rl.DrawRectangleLines(posX, posY, width, height, ctp.Mantle)
	rl.DrawTextEx(font, filepath.Base(d.Texture), rl.Vector2{X: float32(posX), Y: float32(posY)}, 16, 2, ctp.Text)
}
