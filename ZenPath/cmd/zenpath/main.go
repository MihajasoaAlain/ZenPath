package main

import (
	"fmt"
	"os"
	"sync"

	"ZenPath/internal/organizer"
	"ZenPath/internal/scanner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: zenPath <dossier>")
		return
	}

	baseDir := os.Args[1]

	files, err := scanner.Scan(baseDir)
	if err != nil {
		fmt.Println("Erreur scan :", err)
		return
	}

	jobs := make(chan string)
	var wg sync.WaitGroup

	workers := 4

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go organizer.Worker(&wg, jobs, baseDir)
	}

	for _, file := range files {
		jobs <- file
	}

	close(jobs)
	wg.Wait()

	fmt.Println("Organisation terminée ")
}
