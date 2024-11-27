package world

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/bloodmagesoftware/architect/internal/config"
	"github.com/bloodmagesoftware/architect/internal/file"
	"github.com/bloodmagesoftware/architect/internal/version"
	"github.com/charmbracelet/log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

type (
	World struct {
		Version    version.Version `toml:"version"`
		Name       string          `toml:"-"`
		Layers     []Layer         `toml:"layer"`
		Collisions []Collision     `toml:"collision"`
	}

	Layer []Object

	Object struct {
		Texture string `toml:"texture"`
		X       int    `toml:"x"`
		Y       int    `toml:"y"`
	}

	Collision struct {
		X      int `toml:"x"`
		Y      int `toml:"y"`
	}
)

func Worlds() ([]*World, error) {
	worldFiles, err := worldFiles()
	if err != nil {
		return nil, errors.Join(errors.New("failed to get world files"), err)
	}

	var worlds []*World
	for _, worldFile := range worldFiles {
		if !strings.HasSuffix(worldFile, ".toml") {
			continue
		}
		f, err := os.Open(worldFile)
		if err != nil {
			return nil, errors.Join(errors.New("failed to open world file "+worldFile), err)
		}
		defer f.Close()

		world, err := loadWorld(f)
		if err != nil {
			return nil, errors.Join(errors.New("failed to load world "+worldFile), err)
		}
		worlds = append(worlds, world)
	}

	return worlds, nil
}

func randomWorldName() string {
	return fmt.Sprintf("world-%x.toml", rand.Int31())
}

func NewWorld() (*World, error) {
	log.Debug("creating new world")
	worldsDir, err := file.WorldsDir()
	if err != nil {
		return nil, errors.Join(errors.New("failed to get worlds directory"), err)
	}

	p := filepath.Join(worldsDir, randomWorldName())
	f, err := os.Create(p)
	if err != nil {
		return nil, errors.Join(errors.New("failed to create world file "+p), err)
	}
	defer f.Close()

	n := filepath.Base(p)
	w := World{
		Name:    n[:len(n)-5],
		Version: config.Version,
	}

	log.Debug("new world", "w", w)

	e := toml.NewEncoder(f)
	if err := e.Encode(w); err != nil {
		return nil, errors.Join(errors.New("failed to encode world"), err)
	}

	return &w, nil
}

func worldFiles() ([]string, error) {
	worldsDir, err := file.WorldsDir()
	if err != nil {
		return nil, errors.Join(errors.New("failed to get worlds directory"), err)
	}

	dirEntries, err := os.ReadDir(worldsDir)
	if err != nil {
		return nil, errors.Join(errors.New("failed to read directory "+worldsDir), err)
	}

	var worldFiles []string
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			continue
		}
		worldFiles = append(worldFiles, filepath.Join(worldsDir, dirEntry.Name()))
	}
	return worldFiles, nil
}

func loadWorld(f *os.File) (*World, error) {
	w := new(World)
	d := toml.NewDecoder(f)
	_, err := d.Decode(&w)
	if err != nil {
		return nil, errors.Join(errors.New("failed to decode world"), err)
	}
	n := filepath.Base(f.Name())
	w.Name = n[:len(n)-5]
	return w, nil
}
