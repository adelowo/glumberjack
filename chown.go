//go:build !linux
// +build !linux

package glumberjack

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
