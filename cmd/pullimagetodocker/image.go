package main

import (
	"fmt"
	"github.com/ooooo-youwillsee/ooooo-tools/pkg/file"
	"os/exec"
	"strings"
	"sync"
)

func extractImagesForLine(line, imagePrefix string) []string {
	var images []string
	cnt := strings.Count(line, imagePrefix)
	for i := 0; i < cnt; i++ {
		begin := strings.Index(line, imagePrefix)
		for end := begin; end < len(line); end++ {
			// find corresponding symbol or reached to end positon
			if line[end] == line[begin-1] || end == len(line)-1 {
				image := strings.TrimSpace(line[begin:end])
				image = strings.Split(image, "@")[0]
				images = append(images, image)
				line = line[end+1:]
				break
			}
		}
	}
	return images
}

func extractImages() []string {
	if verbose {
		fmt.Println("Extract images:")
		fmt.Println()
	}

	var images []string
	if filePath != "" {
		file.TravelPathEachLine(filePath, func(line string) {
			images = append(images, extractImagesForLine(line, imagePrefix)...)
		})
	}

	if dirPath != "" {
		file.TravelPathEachLine(dirPath, func(line string) {
			images = append(images, extractImagesForLine(line, imagePrefix)...)
		})
	}

	if verbose {
		fmt.Println("images: ")
		for _, image := range images {
			fmt.Println(image)
		}
		fmt.Println()
	}
	return images
}

func pullImages(images []string) {
	if verbose {
		fmt.Println()
		fmt.Println("Pull images:")
		fmt.Println()
	}

	var wg sync.WaitGroup
	for _, image := range images {
		image := image
		wg.Add(1)
		go func() {
			cmd := exec.Command("docker", "pull", image)
			err := cmd.Run()
			if verbose || err != nil {
				fmt.Printf("Exec command '%s', err: %v \n", cmd.String(), err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func tagImages(images []string) []string {
	if verbose {
		fmt.Println()
		fmt.Println("Tag images:")
		fmt.Println()
	}

	var newImages []string
	newTag := username
	for _, image := range images {
		// gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd/controller:v0.47.0@sha256:f2d03e5b00345da4bf91044daff32795f6f54edb23f8a36742abd729929c7943
		newImage := strings.ReplaceAll(image, imagePrefix, newTag)
		if verbose {
			fmt.Printf("Tag %s to %s.\n", image, newImage)
		}
		newImages = append(newImages, newImage)
	}
	return newImages
}

func pushImages(images []string) {
	if verbose {
		fmt.Println()
		fmt.Println("Push images:")
		fmt.Println()
	}

	var wg sync.WaitGroup
	for _, image := range images {
		image := image
		wg.Add(1)
		go func() {
			cmd := exec.Command("docker", "push", image)
			err := cmd.Run()
			if verbose || err != nil {
				fmt.Printf("Exec command '%s', err: %v \n", cmd.String(), err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
