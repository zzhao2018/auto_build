package query_insert_build_tools

type BuildInterface interface{
	collectImport(importElement string)
	Build() (string)
	GetImportList()[]string
}



type WherePair struct {
	Name string
	Opt string   // =  >  >=  <  <=   <>
	Type string
}

func NewWherePair(name string,opt string,types string)*WherePair{
	return &WherePair{
		Name: name,
		Opt: opt,
		Type: types,
	}
}

type QueryPair struct{
	Name string
	Type string
}

func NewQueryPair(name string,types string)*QueryPair{
	return &QueryPair{
		Name: name,
		Type: types,
	}
}