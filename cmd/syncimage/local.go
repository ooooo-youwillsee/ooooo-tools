package main

import (
	"github.com/ooooo-youwillsee/ooooo-tools/pkg/image"
	"github.com/spf13/cobra"
	"os"
)

var (
	output string

	localCmd = cobra.Command{
		Use:  "local",
		Long: "pull image, save image to local",
		Run: func(cmd *cobra.Command, args []string) {
			syncImageToLocal()
		},
	}
)

func init() {
	flags := localCmd.Flags()
	flags.StringVarP(&output, "output", "o", "", "-o output")
	localCmd.MarkFlagRequired("output")

	rootCmd.AddCommand(&localCmd)
}

func syncImageToLocal() {
	var (
		images []string
	)
	if filePath != "" {
		images = image.ExtractImages(filePath, imagePrefix, verbose)
	} else {
		images = image.ExtractImages(dirPath, imagePrefix, verbose)
	}

	os.MkdirAll(output, 0644)
	image.PullImages(images, env, verbose)
	image.SaveImage(images, output, verbose)
}
