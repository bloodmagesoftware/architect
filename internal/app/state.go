package app

import (
	"errors"
	"github.com/bloodmagesoftware/architect/internal/world"
)

var (
	ErrInvalidWorld = errors.New("invalid world")
)

var (
	Worlds []*world.World
	World  int
)

func init() {
	var err error
	Worlds, err = world.Worlds()
	if err != nil {
		panic(err)
	}
	World = -1
}

func SetWorld(i int) error {
	if i < 0 || i >= len(Worlds) {
		return ErrInvalidWorld
	}

	World = i
	return nil
}

func NewWorld() error {
	_, err := world.NewWorld()
	if err != nil {
		return errors.Join(errors.New("failed to create new world"), err)
	}
	Worlds, err = world.Worlds()
	if err != nil {
		return errors.Join(errors.New("failed to update worlds"), err)
	}
	for i, w := range Worlds {
		if w.Name == w.Name {
			World = i
			return nil
		}
	}
	return errors.New("failed to find new world")
}
