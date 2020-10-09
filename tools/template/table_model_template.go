package template

const ConstantClickhouseInterfaceTemplate=`
public interface ClickHouseTable {
    public String getTableName();
}
`

const ConstantClickhouseTableTemplate=`
public class {{className}} implements ClickHouseTable{
	{{fieldName}}
	
	@Override
	public String getTableName() {
        return "{{databaseName}}.{{tableName}}"";
    }
}
`