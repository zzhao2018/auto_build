package query_insert_build_tools

import (
	"fmt"
	"testing"
)

func TestConfigBuildTools_BuildSqlConfig(t *testing.T) {
	queryField:=[]*QueryPair{
		NewQueryPair("ftime","java.util.Date"),
		NewQueryPair("ftrack_id","long"),
	}

	whereField:=[]*WherePair{
		NewWherePair("ftime","=","java.util.Date"),
		NewWherePair("id",">","int"),
	}

	buildTools:=NewQueryConfigBuildTools(queryField,"bi_second.ods_music_song_copyright_info_df",
		whereField,"select %s from %s where {{wherefield}} and pkg>10 and {{wherefield}} and status=1")
	fmt.Println(buildTools.Build())
	fmt.Println(buildTools.GetImportList())
}
