package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/adimail/torrent-client/internal/torrentfile"
	"github.com/manifoldco/promptui"
)

func findTorrentFiles(dir string, torrentFiles *[]string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".torrent") {
			*torrentFiles = append(*torrentFiles, path)
		}
		return nil
	})

	return err
}

func main() {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	var torrentFiles []string
	err = findTorrentFiles(dir, &torrentFiles)
	if err != nil {
		fmt.Println("Error finding torrent files:", err)
		return
	}

	if len(torrentFiles) == 0 {
		fmt.Println("No torrent files detected.")
		return
	} else {
		fmt.Println(len(torrentFiles), "torrent files detected")
		fmt.Println("-----------------------------------------")

	}

	prompt := promptui.Select{
		Label: "Select a file to proceed:",
		Items: torrentFiles,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "\U0001F449 {{ . | cyan }}",
			Inactive: "  {{ . | faint }}",
			Selected: "\U0001F64C {{ . | green | bold }}",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You selected %s\n", result)
	fmt.Printf("Downloading... %s\n", result)

	//  File path read successfully
	inPath := result
	outPath := "./downloads"

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
