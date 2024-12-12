package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	downloadsFolder := filepath.Join(os.Getenv("HOME"), "Downloads")
	err = watcher.Add(downloadsFolder)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	go func() {
		fmt.Printf("Detecting Windows Files \n")
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					if filepath.Ext(event.Name) == ".docx" || filepath.Ext(event.Name) == ".doc" || filepath.Ext(event.Name) == ".pptx" {
						fmt.Printf("File Detected: Begin Convert.\n")
						convertToPDF(event.Name)
					}
				}
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()

	<-done
}

func convertToPDF(docxFile string) {
	pdfFile := docxFile[:len(docxFile)-len(filepath.Ext(docxFile))] + ".pdf"

	cmd := exec.Command("soffice", "--headless", "--convert-to", "pdf", docxFile, "--outdir", filepath.Dir(docxFile))
	err := cmd.Run()
	if err != nil {
		log.Println("Error converting to PDF:", err)
		return
	}

	err = os.Remove(docxFile)
	if err != nil {
		log.Println("Error moving .docx/.doc to trash:", err)
		return
	}

	fmt.Printf("Converted %s to %s See ya. \n", docxFile, pdfFile)
}
