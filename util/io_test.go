package util

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetFiles(t *testing.T) {
	dir := GetGoPath() + ""
	files := GetFilesRecursive(dir)
	fmt.Printf("%v\n", strings.Join(files, "\n"))
}
