package file

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 遍历文件夹
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

// 遍历路径，处理每一行, 适配[文件]和[文件夹]
func TravelPathForEachLine(path string, handle func(line string)) {
	if IsDir(path) {
		TravelDirForEachLine(path, handle)
	} else {
		TravelFileForEachLine(path, handle)
	}
}

// 遍历文件夹，处理每一行
func TravelDirForEachLine(dirPath string, handle func(line string)) {
	TravelDir(dirPath, func(filePath string) {
		TravelFileForEachLine(filePath, handle)
	})
}

// 遍历文件，处理每一行
func TravelFileForEachLine(filePath string, handle func(line string)) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
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

// 复制目录， 处理每一行
func CopyPathForEachLine(path, newPath string, handle func(line string) string) {
	if IsDir(path) && IsDir(newPath) {
		CopyDirForEachLine(path, newPath, handle)
	}
	if !IsDir(path) && !IsDir(newPath) {
		CopyFileForEachLine(path, newPath, handle)
	}
}

// 复制文件夹，处理每一行
func CopyDirForEachLine(dirPath string, newDirPath string, handle func(line string) string) {
	err := os.MkdirAll(newDirPath, 0644)
	if err != nil {
		panic(fmt.Errorf("Mkdir %s, err: %v", newDirPath, err))
	}

	TravelDir(dirPath, func(filePath string) {
		relativePath := strings.TrimPrefix(filePath, dirPath)
		newFilePath := filepath.Join(newDirPath, relativePath)
		CopyFileForEachLine(filePath, newFilePath, handle)
	})
}

// 复制文件，处理每一行
func CopyFileForEachLine(filepath string, newFilePath string, handle func(line string) string) {
	file, err := os.OpenFile(newFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	if err != nil {
		panic(fmt.Errorf("Open File %s, err :%v", newFilePath, err))
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	TravelFileForEachLine(filepath, func(line string) {
		newLine := handle(line)
		writer.WriteString(newLine)
	})
}

func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		panic(err)
	}
	return stat.IsDir()
}
