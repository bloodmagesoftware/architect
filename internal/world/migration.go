package world

import (
	"fmt"
	"github.com/bloodmagesoftware/architect/internal/config"
)

func (w *World) Migrate() error {
	switch w.Version.Compare(config.Version) {
	case 1:
		if config.Upgrade {
			return w.doMigration()
		} else {
			return fmt.Errorf("world version (%s) is newer than architect version (%s), update architect to use this world file", w.Version.String(), config.Version.String())
		}
	case 0:
		return nil
	case -1:
		return fmt.Errorf("world version (%s) is older than architect version (%s), update the world file using the --upgrade flag to use this world file", w.Version.String(), config.Version.String())
	default:
		return fmt.Errorf("unknown error comparing world version (%s) to architect version (%s)", w.Version.String(), config.Version.String())
	}
}

func (w *World) doMigration() error {
	return fmt.Errorf("migration from %s to %s not implemented", w.Version.String(), config.Version.String())
}
