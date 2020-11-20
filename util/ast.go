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
