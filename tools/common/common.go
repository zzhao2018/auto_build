package common

import (
	"strings"
)

func IsDigital(typeName string)(bool){
	if strings.Contains(strings.ToLower(typeName),"int")||
		strings.Contains(strings.ToLower(typeName),"long"){
		return true
	}
	return false
}

func split_format(r rune)bool{
	splitArr:=[]rune{'.','_'}
	for _,splitEle:=range splitArr {
		if r==splitEle {
			return true
		}
	}
	return false
}

func LineToCamelCase(tableName string)string{
	var (
		tableNameElementArr=make([]string,0)
		loc=0
	)
	tableNameElementArr=strings.FieldsFunc(tableName,split_format)
	for _,tableNameEle:=range tableNameElementArr {
		tableNameArr:=[]byte(tableNameEle)
		if tableNameArr!=nil&&len(tableNameArr)>0&&tableNameArr[0]>='a' && tableNameArr[0]<='z'{
			tableNameArr[0]=tableNameArr[0]-'a'+'A'
		}
		tableNameElementArr[loc]=string(tableNameArr)
		loc++
	}
	return strings.Join(tableNameElementArr,"")
}

func ReplaceArr(sourceStr string,oldStrArr []string,newStrArr []string)string{
	if len(oldStrArr)!=len(newStrArr) {
		return sourceStr
	}
	for loc:=0;loc<len(oldStrArr);loc++ {
		sourceStr=strings.Replace(sourceStr,oldStrArr[loc],newStrArr[loc],1)
	}
	return sourceStr
}