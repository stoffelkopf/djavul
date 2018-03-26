// Package assets provides utility functions for handling game assets.
package assets

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// GameDir returns the path of the game directory.
func GameDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", errors.WithStack(err)
	}
	gameDir := filepath.Dir(exePath)
	return gameDir, nil
}
