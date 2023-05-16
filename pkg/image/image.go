package image

import (
	"fmt"
	"github.com/ooooo-youwillsee/ooooo-tools/pkg/file"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func extractImagesForLine(line, imagePrefix string) []string {
	var images []string
	cnt := strings.Count(line, imagePrefix)
	for i := 0; i < cnt; i++ {
		begin := strings.Index(line, imagePrefix)
		for end := begin; end < len(line); end++ {
			// find corresponding symbol or reached to end positon
			if end == len(line)-1 {
				image := strings.TrimSpace(line[begin:])
				images = append(images, image)
				break
			}
			if line[end] == line[begin-1] {
				image := strings.TrimSpace(line[begin:end])
				images = append(images, image)
				line = line[end:]
				break
			}
		}
	}
	return images
}

func ExtractImages(path, imagePrefix string, verbose bool) []string {
	if verbose {
		fmt.Println("Extract images:")
		fmt.Println()
	}

	var images []string
	file.TravelPathForEachLine(path, func(line string) {
		images = append(images, extractImagesForLine(line, imagePrefix)...)
	})

	if verbose {
		for _, image := range images {
			fmt.Println(image)
		}
		fmt.Println()
	}
	return images
}

func PullImages(images []string, env []string, verbose bool) {
	if verbose {
		fmt.Println("Pull images:")
		fmt.Println()
	}

	for _, image := range images {
		cmd := exec.Command("docker", "pull", image)
		cmd.Env = append(os.Environ(), env...)
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

func RenameImage(image, repository string) string {
	// gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd/controller:v0.47.0@sha256:f2d03e5b00345da4bf91044daff32795f6f54edb23f8a36742abd729929c7943
	image = strings.Split(image, "@")[0]
	segs := strings.Split(image, ":")
	imageUrl, version := segs[0], segs[1]

	segs = strings.Split(imageUrl, "/")
	if strings.Contains(segs[0], ".") {
		segs[0] = strings.ReplaceAll(segs[0], ".", "_")
	}
	imageName := strings.Join(segs, "_")

	image = fmt.Sprintf("%s/%s:%s", repository, imageName, version)
	return image
}

func RenameImages(images []string, repository string, verbose bool) []string {
	if verbose {
		fmt.Println("Rename Images:")
		fmt.Println()
	}
	var newImages []string
	for _, image := range images {
		newImage := RenameImage(image, repository)
		newImages = append(newImages, newImage)
		if verbose {
			fmt.Printf("Rename %s to %s.\n", image, newImage)
		}
	}
	fmt.Println()
	return newImages
}

func TagImages(images, newImages []string, verbose bool) {
	if verbose {
		fmt.Println("Tag images:")
		fmt.Println()
	}
	for i := range images {
		image, newImage := images[i], newImages[i]
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
	}
	fmt.Println()
}

func PushImages(images []string, verbose bool) {
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

func SaveImage(images []string, output string, verbose bool) {
	for _, image := range images {
		image = strings.Split(image, "@")[0]
		tarName := strings.ReplaceAll(strings.ReplaceAll(image, "/", "_"), ":", "__")
		tarFile := filepath.Join(output, tarName+".tar")
		cmd := exec.Command("docker", "save", "-o", tarFile, image)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if verbose {
			fmt.Println(cmd.String())
		}
		err := cmd.Run()
		if err != nil {
			panic(fmt.Errorf("Save Image, err: %v", err))
		}
	}
}
