package lib

import (
	"os"

	ozlog "github.com/xqm123/oozlog/go"
)

// readJsonAndGen 从文件中读取json信息并且生成struct
func ReadJsonAndGen(jsonFile, outType, outFile string) {
	if len(jsonFile) == 0 {
		jsonFile = DefaultJsonFile
	}
	file, err := os.OpenFile(jsonFile, os.O_RDONLY, 0666)
	if err != nil {
		ozlog.Errorf("open file err")
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		ozlog.Errorf("read file stat err")
	}
	filesize := fileinfo.Size()

	var (
		json string
	)
	buf := make([]byte, filesize)

	len, _ := file.Read(buf)
	if len == 0 {
		return
	}
	json = `` + string(buf[:len]) + ``
	ck := New(json, jsonFile, outType, outFile)
	ck.json2Struct()

	return
}
