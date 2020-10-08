package query_insert_build_tools

import (
	"auto_build/tools/common"
	"auto_build/tools/template"
	"fmt"
	"strings"
)

type NameTypePair struct {
	Name string
	Opt string   // =  >  >=  <  <=   <>
	Type string
}

type ConfigBuildTools struct {
	queryField []NameTypePair
	fromTableName string
	whereField []NameTypePair
	querySql string
	//`select {{queryFields}} from {{queryTable}} where {{wherefield}} and {{wherefiled}} and {{wherefiled}} ...`
}


func NewBuildTools(queryField []NameTypePair,fromTableName string,whereField []NameTypePair,querySql string)(*ConfigBuildTools){
	return &ConfigBuildTools{
		queryField: queryField,
		fromTableName: fromTableName,
		whereField: whereField,
		querySql: querySql,
	}
}

func (c *ConfigBuildTools)BuildSqlConfig() (string) {
	var (
		loc=0
		paramArr=make([]string, len(c.whereField))
		returnFieldArr=make([]string,len(c.whereField))
	)
	//1. 根据wherefield生成查询参数与where参数
	for _,whereFieldPair:=range c.whereField{
		paramArr[loc]=fmt.Sprintf("%s %s",whereFieldPair.Name,whereFieldPair.Type)
		if common.IsDigital(whereFieldPair.Type)==true{
			returnFieldArr[loc]=fmt.Sprint("%s%s+\"+%s+\"",whereFieldPair.Name,whereFieldPair.Opt,whereFieldPair.Name)
		}else{
			returnFieldArr[loc]=fmt.Sprint("%s%s'\"+%s+\"'",whereFieldPair.Name,whereFieldPair.Opt,whereFieldPair.Name)
		}
		loc++
	}

	QuerySqlSplitByWhere:=strings.Split(c.querySql,"{{wherefield}}")

	QeuryName:=common.LineToCamelCase(c.fromTableName)
	QuerySql:=QuerySqlSplitByWhere[0]
	QueryField:=strings.Join(paramArr,",")
	QueryReturnSql:=QeuryName

	for _,QuerySqlField:=range QuerySqlSplitByWhere[1:]{
		QueryReturnSql=QueryReturnSql+"+"+returnFieldArr[loc]+"+"+QuerySqlField;
	}



	return strings.ReplaceAll(
		strings.ReplaceAll(template.ConstantQuerySqlTemplate,"{{QeuryName}}",QeuryName),"{{QuerySql}}",QuerySql)
}

