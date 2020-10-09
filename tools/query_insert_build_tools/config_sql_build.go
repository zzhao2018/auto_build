package query_insert_build_tools

import (
	"auto_build/tools/common"
	"auto_build/tools/template"
	"fmt"
	"strings"
)


// 查询sql与方法生成
type QueryConfigBuildTools struct {
	queryField []*QueryPair
	fromTableName string
	whereField []*WherePair
	querySql string
	importCollect []string
}


func NewQueryConfigBuildTools(queryField []*QueryPair,fromTableName string,whereField []*WherePair,querySql string)(*QueryConfigBuildTools){
	return &QueryConfigBuildTools{
		queryField: queryField,
		fromTableName: fromTableName,
		whereField: whereField,
		querySql: querySql,
		importCollect: make([]string,0),
	}
}

func(c *QueryConfigBuildTools)collectImport(importElement string){
	c.importCollect=append(c.importCollect,importElement)
}


func (c *QueryConfigBuildTools)Build() (string) {
	paramArr,returnFieldArr:=c.genParam()
	return c.genCode(paramArr,returnFieldArr)
}

func(c *QueryConfigBuildTools)genParam()([]string,[]string){
	var (
		loc=0
		paramArr=make([]string, len(c.whereField))
		returnFieldArr=make([]string,len(c.whereField))
	)
	//1. 根据wherefield生成查询参数与where参数
	for _,whereFieldPair:=range c.whereField{
		importType,shouldImport:=common.JuageImportType(whereFieldPair.Type)
		if shouldImport==true {
			c.collectImport(importType)
		}
		paramArr[loc]=fmt.Sprintf("%s %s",whereFieldPair.Type,whereFieldPair.Name)
		if common.IsDigital(whereFieldPair.Type)==true{
			returnFieldArr[loc]=fmt.Sprintf("\"%s%s\"+%s",whereFieldPair.Name,whereFieldPair.Opt,whereFieldPair.Name)
		}else{
			returnFieldArr[loc]=fmt.Sprintf("\"%s%s'\"+%s+\"'\"",whereFieldPair.Name,whereFieldPair.Opt,whereFieldPair.Name)
		}
		loc++
	}
	return paramArr,returnFieldArr
}

func(c *QueryConfigBuildTools)genCode(paramArr []string,returnFieldArr []string)(string){
	var(
		loc=0
		replaceTag=[]string{"{{QeuryName}}","{{QuerySql}}","{{QueryField}}","{{QueryReturnSql}}"}
		replaceVal=[]string{}
		QueryReturnFieldArr=[]string{}
	)

	QueryName:=common.LineToCamelCase(c.fromTableName)
	QueryField:=strings.Join(paramArr,",")
	QuerySqlSplitByWhere:=strings.Split(c.querySql,"{{wherefield}}")
	QuerySql:="\""+QuerySqlSplitByWhere[0]+"\""
	QueryReturnFieldArr=append(QueryReturnFieldArr,QueryName)
	for _,QuerySqlField:=range QuerySqlSplitByWhere[1:]{
		QueryReturnFieldArr=append(QueryReturnFieldArr,returnFieldArr[loc],"\""+QuerySqlField+"\"")
		loc++
	}
	QueryReturnSql:=strings.Join(QueryReturnFieldArr,"+")
	//整理生成的sql拼接语句
	for{
		CopyReturnSql:=QueryReturnSql
		QueryReturnSql=common.ReplaceArr(QueryReturnSql,[]string{"\"+\"","+\"\""},[]string{"",""})
		if CopyReturnSql==QueryReturnSql {
			break
		}
	}
	replaceVal = append(replaceVal, QueryName,QuerySql,QueryField,QueryReturnSql)
	return common.ReplaceArr(template.ConstantQuerySqlTemplate,replaceTag,replaceVal)
}


func(c *QueryConfigBuildTools)GetImportList()[]string{
	return c.importCollect
}




//table生成
type QueryTableBuildTools struct{
	
}

func NewQueryTableBuildTools() *QueryTableBuildTools {
	return &QueryTableBuildTools{}
}


