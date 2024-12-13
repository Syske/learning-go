package main

import (
	"fmt"

	"syske.github.io/test/go_test/learning"
	"syske.github.io/test/pkg1"
	"syske.github.io/test/util/db_util"
	"syske.github.io/test/util/net_util"
)

func main() {

	db_util.TestDb()
	// db_util.TestPdb()
	learning.Func_args("1213", "56756", "67867")
	learning.CallBackTest()
	pkg1.PkgTest()
	pkg1.ForInStr("hello")
	url := "https://coolapi.coolcollege.cn/exam-api/v2/enterprises/1174579893367869459/export_record/composeCount?access_token=8a32a847c444487f867022df8f28e7ce"
	net_util.GetRequest(url)
	contentType := "application/json"
	data := "{}"
	result := net_util.PostRequest(url, contentType, data)
	fmt.Println(result)

}
