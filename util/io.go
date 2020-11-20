// Copyright (c) 2012-2020 Grabtaxi Holdings PTE LTD (GRAB), All Rights Reserved. NOTICE: All information contained herein
// is, and remains the property of GRAB. The intellectual and technical concepts contained herein are confidential, proprietary
// and controlled by GRAB and may be covered by patents, patents in process, and are protected by trade secret or copyright law.
//
// You are strictly forbidden to copy, download, store (in any medium), transmit, disseminate, adapt or change this material
// in any way unless prior written permission is obtained from GRAB. Access to the source code contained herein is hereby
// forbidden to anyone except current GRAB employees or contractors with binding Confidentiality and Non-disclosure agreements
// explicitly covering such access.
//
// The copyright notice above does not evidence any actual or intended publication or disclosure of this source code,
// which includes information that is confidential and/or proprietary, and is a trade secret, of GRAB.
//
// ANY REPRODUCTION, MODIFICATION, DISTRIBUTION, PUBLIC PERFORMANCE, OR PUBLIC DISPLAY OF OR THROUGH USE OF THIS SOURCE
// CODE WITHOUT THE EXPRESS WRITTEN CONSENT OF GRAB IS STRICTLY PROHIBITED, AND IN VIOLATION OF APPLICABLE LAWS AND
// INTERNATIONAL TREATIES. THE RECEIPT OR POSSESSION OF THIS SOURCE CODE AND/OR RELATED INFORMATION DOES NOT CONVEY
// OR IMPLY ANY RIGHTS TO REPRODUCE, DISCLOSE OR DISTRIBUTE ITS CONTENTS, OR TO MANUFACTURE, USE, OR SELL ANYTHING
// THAT IT MAY DESCRIBE, IN WHOLE OR IN PART.

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
