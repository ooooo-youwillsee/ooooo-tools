package main

import (
	"fmt"
	"github.com/ooooo-youwillsee/ooooo-tools/pkg/file"
	"os"
	"os/exec"
	"strings"
)

func ExtractImagesForLine(line, imagePrefix string) []string {
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

func ExtractImages() []string {
	if verbose {
		fmt.Println("Extract images:")
		fmt.Println()
	}

	var images []string
	if filePath != "" {
		file.TravelPathEachLine(filePath, func(line string) {
			images = append(images, ExtractImagesForLine(line, imagePrefix)...)
		})
	}

	if dirPath != "" {
		file.TravelPathEachLine(dirPath, func(line string) {
			images = append(images, ExtractImagesForLine(line, imagePrefix)...)
		})
	}

	if verbose {
		for _, image := range images {
			fmt.Println(image)
		}
		fmt.Println()
	}
	return images
}

func PullImages(images []string) {
	if verbose {
		fmt.Println()
		fmt.Println("Pull images:")
		fmt.Println()
	}

	for _, image := range images {
		cmd := exec.Command("docker", "pull", image)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if verbose {
			fmt.Println(cmd.String())
		}
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		fmt.Println()
	}
}

func RenameImage(image, username string) string {
	// gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd/controller:v0.47.0@sha256:f2d03e5b00345da4bf91044daff32795f6f54edb23f8a36742abd729929c7943
	image = strings.Split(image, "@")[0]
	segs := strings.Split(image, ":")
	imageUrl, version := segs[0], segs[1]

	segs = strings.Split(imageUrl, "/")
	if strings.Contains(segs[0], ".") {
		segs[0] = strings.ReplaceAll(segs[0], ".", "_")
	}
	imageName := strings.Join(segs, "_")

	image = fmt.Sprintf("%s/%s/%s:%s", "docker.io", username, imageName, version)
	return image
}

func TagImages(images []string) []string {
	if verbose {
		fmt.Println("Tag images:")
		fmt.Println()
	}

	var newImages []string
	for _, image := range images {
		newImage := RenameImage(image, username)
		cmd := exec.Command("docker", "tag", image, newImage)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if verbose {
			fmt.Printf("Tag %s to %s.\n", image, newImage)
		}
		err := cmd.Run()
		if err != nil {
			panic(err)
		}

		newImages = append(newImages, newImage)
	}
	return newImages
}

func pushImages(images []string) {
	if verbose {
		fmt.Println("Push images:")
		fmt.Println()
	}

	for _, image := range images {
		cmd := exec.Command("docker", "push", image)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if verbose {
			fmt.Println(cmd.String())
		}
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		fmt.Println()
	}
}
