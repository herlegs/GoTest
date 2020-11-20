package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/herlegs/GoTest/util"
)

// DtoInfo ...
type DtoInfo struct {
	DtoName             string
	STT                 *ast.StructType
	FoundInFile         string
	PartitionField      string
	PartitionIDFuncBody string
}

// GetDTOInfo return all parsed information
func GetDTOInfo(streamFolder string) (dtoInfo *DtoInfo) {
	dtoInfo = &DtoInfo{}

	goFiles := util.GetGoFileNames(streamFolder)

	foundDto := false

	for _, filename := range goFiles {
		file, err := parser.ParseFile(token.NewFileSet(), streamFolder+"/"+filename, nil, parser.ParseComments)
		if err != nil {
			fmt.Printf("err parsing file %v: %v\n", filename, err)
			return
		}

		ast.Inspect(file, func(node ast.Node) bool {
			// check for struct def
			typeSpec, ok := node.(*ast.TypeSpec)
			if !ok {
				return true
			}
			if stt, ok := typeSpec.Type.(*ast.StructType); ok {
				fields := stt.Fields.List

				for _, field := range fields {
					if len(field.Names) > 0 {
						fieldName := field.Names[0].Name
						fmt.Printf("%v\n", fieldName)
					}
				}
			}
			return foundDto
		})
	}

	if foundDto {
		fmt.Printf("Found stream dto [%v] in file [%v]\n", dtoInfo.DtoName, dtoInfo.FoundInFile)

		// get fields for converter generation
		_, _ = parser.ParseFile(token.NewFileSet(), streamFolder+"/"+dtoInfo.FoundInFile, nil, parser.ParseComments)

		// guess partition ID, and function migration
		for _, filename := range goFiles {
			fileContent := util.ReadFile(streamFolder + "/" + filename)

			// migrate GetPartitionID function
			fs := token.NewFileSet()
			file, _ := parser.ParseFile(fs, "", fileContent, parser.ParseComments)
			ast.Inspect(file, func(node ast.Node) bool {
				fc, ok := node.(*ast.FuncDecl)
				if ok && fc.Name.Name == "TargetFunc" && fc.Recv == nil {
					for _, param := range fc.Type.Params.List {
						paramName, paramType := param.Names[0].String(), util.GetNodeString(fileContent, fs, param.Type)
						funcBody := util.GetNodeString(fileContent, fs, fc.Body)

						_, _, _ = paramName, paramType, funcBody
					}
					return false
				}
				return true
			})
		}
	}

	return
}
