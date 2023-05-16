package main

import (
	"fmt"
	"github.com/ooooo-youwillsee/ooooo-tools/pkg/file"
	"github.com/ooooo-youwillsee/ooooo-tools/pkg/image"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	repository string

	remoteCmd = cobra.Command{
		Use:  "remote",
		Long: "pull image, retag image, push imageRepository",
		Run: func(cmd *cobra.Command, args []string) {
			syncImageToRemote()
		},
	}
)

func init() {
	flags := remoteCmd.Flags()
	flags.StringVarP(&repository, "repository", "r", "", "-r docker.io/youwillsee")
	remoteCmd.MarkFlagRequired("repository")

	rootCmd.AddCommand(&remoteCmd)
}

func syncImageToRemote() {
	var (
		images    []string
		newImages []string
		path      string
		newPath   string
	)
	if filePath != "" {
		images = image.ExtractImages(filePath, imagePrefix, verbose)
		path = filePath
		newPath = filePath + ".bak"
		os.Create(newPath)
	} else {
		images = image.ExtractImages(dirPath, imagePrefix, verbose)
		path = dirPath
		newPath = filePath + ".bak"
		os.MkdirAll(newPath, 0644)
	}

	image.PullImages(images, env, verbose)
	newImages = image.RenameImages(images, repository, verbose)
	image.TagImages(images, newImages, verbose)
	image.PushImages(newImages, verbose)
	generatorNewK8sFile(path, newPath, images, newImages)
}

func generatorNewK8sFile(path string, newPath string, images []string, newImages []string) {
	file.CopyPathForEachLine(path, newPath, func(line string) string {
		for i := range images {
			image, newImage := images[i], newImages[i]
			line = strings.ReplaceAll(line, image, newImage)
		}
		return line
	})
	fmt.Printf("generator new path %s\n", newPath)
}
