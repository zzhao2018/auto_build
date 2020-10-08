package common

import (
	"fmt"
	"testing"
)

func TestLineToCamelCase(t *testing.T) {
	datas:="bi_sec.ods_music_test_qq"

	fmt.Println(LineToCamelCase(datas))
}