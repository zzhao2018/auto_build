package common

import (
	"auto_build/tools/template"
	"strings"
)

func JuageImportType(typeName string)(string,bool){
	typeName=strings.ToLower(typeName)
	if strings.Contains(typeName,"int")||strings.Contains(typeName,"string") {
		return "",false
	}
	for keys,vals:=range template.ImportMap {
		if keys==typeName {
			return vals,true
		}
	}
	return "",false
}