package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	filePath    = ""
	dirPath     = ""
	imagePrefix = ""
	username    = ""
	output      = ""
	verbose     = false
)

func main() {
	rootCmd := cobra.Command{
		Use:   "pullimagetodocker",
		Short: "pull image to docker",
		Long:  "parse k8s.yaml filePath to extract image url, then docker pull image, tag image to docker.username, push image to docker image registry.",
		Run: func(cmd *cobra.Command, args []string) {
			pullImageToDocker()
		},
	}
	// setting flags
	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&filePath, "filePath", "f", "", "-f k8s.yaml.filePath")
	flags.StringVarP(&dirPath, "dirPath", "d", "", "-d k8s.yaml.dirPath")
	flags.StringVarP(&username, "username", "u", "", "-u docker.username")
	flags.StringVarP(&output, "output", "o", "output", "-o output")
	flags.StringVarP(&imagePrefix, "imagePrefix", "i", "", "-i gcr.io")
	flags.BoolVarP(&verbose, "verbose", "v", false, "-v")
	// check flags
	rootCmd.MarkFlagsMutuallyExclusive("filePath", "dirPath")
	rootCmd.MarkFlagRequired("username")
	rootCmd.MarkFlagRequired("imagePrefix")

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func pullImageToDocker() {
	images := ExtractImages()
	PullImages(images)
	newImages := TagImages(images)
	pushImages(newImages)
}
