package file

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func TravelDir(dirPath string, handle func(filePath string)) {
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		panic(fmt.Sprintf("TravelDir %v", err))
	}

	for _, entry := range dirEntries {
		path := filepath.Join(dirPath, entry.Name())
		if entry.IsDir() {
			TravelDir(path, handle)
		} else {
			handle(path)
		}
	}
}

func TravelPathEachLine(path string, handle func(line string)) {
	stat, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		panic(fmt.Sprintf("path %s is not exist.", path))
	}
	if stat.IsDir() {
		TravelDirEachLine(path, handle)
	} else {
		TravelFileEachLine(path, handle)
	}
}

func TravelDirEachLine(dirPath string, handle func(line string)) {
	TravelDir(dirPath, func(filePath string) {
		TravelFileEachLine(filePath, handle)
	})
}

func TravelFileEachLine(filePath string, handle func(line string)) {
	f, err := os.OpenFile(filePath, os.O_RDWR, 0600)
	defer f.Close()
	if err != nil {
		panic(fmt.Sprintf("OpenFile error %v", err))
	}

	reader := bufio.NewReader(f)
	for true {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		handle(line)
	}
}
