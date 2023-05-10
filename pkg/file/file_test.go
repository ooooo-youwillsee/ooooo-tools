package file

import (
	"fmt"
	"os"
	"testing"
)

func TestTravelDir(t *testing.T) {
	dir, _ := os.Getwd()
	handle := func(path string) {
		fmt.Println(path)
	}

	TravelDir(dir, handle)
}
