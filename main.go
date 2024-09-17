package main

import (
	"github.com/bloodmagesoftware/architect/internal/ui"
	"github.com/bloodmagesoftware/architect/internal/world"
)

func main() {
	worlds, err := world.Worlds()
	if err != nil {
		panic(err)
	}
	for _, w := range worlds {
		err = w.Migrate()
		if err != nil {
			panic(err)
		}
	}

	ui.MainLoop()
}
