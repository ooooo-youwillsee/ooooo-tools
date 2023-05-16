package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	filePath    string
	dirPath     string
	imagePrefix string
	verbose     bool
	env         []string
	processor   []string

	rootCmd = cobra.Command{
		Use:   "syncimage",
		Args:  cobra.MinimumNArgs(1),
		Short: "sync image local or remote",
		Long:  "parse k8s.yaml filePath to extract image url, then docker pull image, save Image to local or push image to remote",
	}
)

func init() {
	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&filePath, "filePath", "f", "", "-f k8s.yaml.filePath")
	flags.StringVarP(&dirPath, "dirPath", "d", "", "-d k8s.yaml.dirPath")
	flags.StringVarP(&imagePrefix, "imagePrefix", "i", "", "-i gcr.io")
	flags.StringSliceVarP(&env, "environment", "e", nil, "-e a=1 -e b=2")
	flags.BoolVarP(&verbose, "verbose", "v", false, "-v")
	flags.StringSliceVarP(&processor, "processor", "", nil, "--processor pullImage,")
	rootCmd.MarkFlagsMutuallyExclusive("filePath", "dirPath")
	rootCmd.MarkPersistentFlagRequired("imagePrefix")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
