package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

func TravelDir(dir string, handle func(path string)) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		panic(fmt.Sprintf("TravelDir %v", err))
	}

	for _, entry := range dirEntries {
		path := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			TravelDir(path, handle)
		} else {
			handle(path)
		}
	}
}
