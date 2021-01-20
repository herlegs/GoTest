package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// ReadFile reads file given full absolute path
func ReadFile(filepath string) string {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("error reading file %v : %v\n", filepath, err)
	}
	return string(bytes)
}

// WriteFile writes to file given full absolute path
func WriteFile(filepath, content string) {
	err := ioutil.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("error writing file %v: %v\n", filepath, err)
	}
}

// FormatGoFile formats file (go imports ) given full absolute path
func FormatGoFile(filepath string) {
	goImportsCmd := GetGoPath() + "/bin/goimports"

	command := exec.Command(goImportsCmd, "-w", "-local", "", filepath)
	if _, err := command.CombinedOutput(); err != nil {
		fmt.Printf("error running goimports : %v", err)
		return
	}
}

// IsFileExist checks if a file exists or not
func IsFileExist(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		return !info.IsDir()
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// IsDirExist checks if a directory exists or not
func IsDirExist(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir()
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// GetFilesRecursive returns all files (path) under a dir
func GetFilesRecursive(dir string) []string {
	var filenames []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("err getting stream folders for dir %v:%v\n", dir, err)
		return filenames
	}
	for _, file := range files {
		filepath := dir + "/" + file.Name()
		if file.IsDir() {
			filenames = append(filenames, GetFilesRecursive(filepath)...)
		} else {
			filenames = append(filenames, filepath)
		}
	}
	return filenames
}
