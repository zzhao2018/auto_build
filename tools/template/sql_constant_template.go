package template

const ConstantFrameTemplate=`
public static class SqlConst{
  {{Content}}
}
`

const ConstantQuerySqlTemplate=`
private static String {{QeuryName}}={{QuerySql}}
public static String gen{{QeuryName}}({{QueryField}}){
   return {{QueryReturnSql}};
}
`

const ConstantInsertSqlTemplate = `
private static String {{InsertName}}={{InsertSql}};
public static String gen{{InsertName}}(){
   return {{InsertReturnSql}};
}
`

