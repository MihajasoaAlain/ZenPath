package organizer

import (
	"path/filepath"
	"strings"

	"ZenPath/internal/config"
)

func DetectFolder(file string) string {
	ext := strings.ToLower(filepath.Ext(file))
	folder, ok := config.Categories[ext]
	if !ok {
		return "Autres"
	}
	return folder
}
