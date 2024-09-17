package file

import (
	"errors"
	"os"
	"path/filepath"
)

func WorldsDir() (string, error) {
	assetsDir, err := assetsDir()
	if err != nil {
		return "", errors.Join(errors.New("failed to get assets directory"), err)
	}
	worldsDir := filepath.Join(assetsDir, "worlds")
	stat, err := os.Stat(worldsDir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(worldsDir, 0755); err != nil {
				return "", errors.Join(errors.New("failed to create directory "+worldsDir), err)
			} else {
				return worldsDir, nil
			}
		} else {
			return "", errors.Join(errors.New("failed to stat "+worldsDir), err)
		}
	}
	if !stat.IsDir() {
		return "", errors.New(worldsDir + " is not a directory")
	}
	return worldsDir, nil
}

func assetsDir() (string, error) {
	stat, err := os.Stat("./assets")
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll("./assets", 0755); err != nil {
				return "", errors.Join(errors.New("failed to create directory ./assets"), err)
			} else {
				return "./assets", nil
			}
		} else {
			return "", errors.Join(errors.New("failed to stat ./assets"), err)
		}
	}
	if !stat.IsDir() {
		return "", errors.New("./assets is not a directory")
	}
	return "./assets", nil
}
