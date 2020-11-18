package getDirTree

import (
	"fmt"
	"io/ioutil"
	"log"
)

var excelRow int = 1
var excelCol int = 0
var excelColAll [] string
var excelSilce = make([]map[string]string, 0,100)
var forFindType string = "1"
//文件目录树形结构节点


func DirTree (dirPath string)([]map[string]string){
	file, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Read dir '%s' failed: %v", file, err)
	}

	for _, fileVal := range file{

		fileName := fileVal.Name()
		fileNamePath := dirPath + "/" + fileName
		//fmt.Println(fileName)
		excelOne := fmt.Sprintf("%s%d",excelColAll[excelCol],excelRow)

		key := string(excelOne)
		//fmt.Printf("%c=%s",key,key)
		dir := make(map[string]string)
		dir[key] = fileName
		excelSilce = append(excelSilce,dir)

		if forFindType == "1"{
			if fileVal.IsDir() {
				excelCol++
				excelRow++
				DirTree(fileNamePath)
				excelCol--
			}
		}

		excelRow++
		//fmt.Println(excelSilce)
	}

	return excelSilce
}

func InitDirTree(dirPath string, findType string)(excelSilce[] map[string]string,err error){

	forFindType = findType
	var a = 'A'
	//生成26个字符
	for i := 1; i <= 26; i++ {
		b := string(a)
		excelColAll = append(excelColAll, b)
		a++
	}

	var aa = 'A'
	for i := 1; i <= 26; i++ {
		b := string(aa)
		b = b + b
		excelColAll = append(excelColAll, b)
		aa++
	}
	re := DirTree(dirPath)
	//fmt.Println(res)
	return re,err

	//fmt.Println("111111")
}
