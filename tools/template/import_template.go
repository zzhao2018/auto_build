package template

var ImportMap = make(map[string]string)

func init()  {
	ImportMap["java.util.date"]=`import java.util.Date;`
	ImportMap["java.sql.date"]=`import java.sql.Date;`
}