package util

import (
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"
)

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"firstUpper": strings.Title,
		"hasPrefix":  strings.HasPrefix,
		"hasSuffix":  strings.HasSuffix,
		"replace":    strings.Replace,
		"toLower":    strings.ToLower,
		"toUpper":    strings.ToUpper,
	}
}

// ReadTemplate ...
func ReadTemplate(tmplPath string, additionalFuncs ...map[string]interface{}) *template.Template {
	tmplContent, err := ioutil.ReadFile(tmplPath)
	if err != nil {
		fmt.Printf("err reading tmpl:%v\n", err)
		return nil
	}

	funcMap := getFuncMap()

	if len(additionalFuncs) > 0 {
		for s, f := range additionalFuncs[0] {
			funcMap[s] = f
		}
	}

	return template.Must(template.New(tmplPath).Funcs(funcMap).Parse(string(tmplContent)))
}
