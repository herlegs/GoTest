package util

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/token"
	"io/ioutil"
	"log"
)

// GetGoFileNames gets all go files (file name only) under directory
func GetGoFileNames(currAbsPath string) []string {
	pkg, err := build.Default.ImportDir(currAbsPath, 0)
	if err != nil {
		log.Fatalf("cannot process current directory %s: %s", currAbsPath, err)
		return nil
	}
	var names []string
	names = append(append(names, pkg.GoFiles...), pkg.TestGoFiles...)
	return names
}

// GetGoFilesRecursive gets all go files (file absolute path) under directory
func GetGoFilesRecursive(currAbsPath string) []string {
	var goFiles []string
	goFileNames := GetGoFileNames(currAbsPath)
	for _, filename := range goFileNames {
		goFiles = append(goFiles, currAbsPath+"/"+filename)
	}

	files, err := ioutil.ReadDir(currAbsPath)
	if err != nil {
		fmt.Printf("err getting stream folders for dir %v:%v\n", currAbsPath, err)
		return goFiles
	}
	for _, file := range files {
		if file.IsDir() {
			goFiles = append(goFiles, GetGoFilesRecursive(currAbsPath+"/"+file.Name())...)
		}
	}
	return goFiles
}

// GetTypeString ...
func GetTypeString(expr ast.Expr) string {
	var result string

	switch etype := expr.(type) {
	case *ast.ArrayType:
		result = fmt.Sprintf("[]%s", GetTypeString(etype.Elt))
	case *ast.MapType:
		result = fmt.Sprintf("map[%s]%s", etype.Key, etype.Value)

	case *ast.SelectorExpr:
		result = fmt.Sprintf("%s.%s", etype.X, etype.Sel)

	case *ast.StarExpr:
		result = fmt.Sprintf("*%s", GetTypeString(etype.X))

	default:
		result = fmt.Sprintf("%s", etype)
	}
	return result
}

// GetNodeString ...
func GetNodeString(fileContent string, fs *token.FileSet, node ast.Node) string {
	start, end := fs.Position(node.Pos()).Offset, fs.Position(node.End()).Offset
	return fileContent[start:end]
}
