package organizer

import (
	"os"
	"path/filepath"
	"sync"
)

func createDir(base, name string) error {
	return os.MkdirAll(filepath.Join(base, name), os.ModePerm)
}

func moveFile(src, dest string) error {
	return os.Rename(src, dest)
}

func Worker(wg *sync.WaitGroup, jobs <-chan string, baseDir string) {
	defer wg.Done()

	for file := range jobs {
		folder := DetectFolder(file)
		destDir := filepath.Join(baseDir, folder)
		createDir(baseDir, folder)
		moveFile(file, filepath.Join(destDir, filepath.Base(file)))
	}
}
