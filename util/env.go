package util

import "go/build"

// GetGoPath ...
func GetGoPath() string {
	return build.Default.GOPATH
}
