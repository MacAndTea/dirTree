package main

import (
	"dirTree/getDirTree"
	"dirTree/mytool"
	"fmt"
	"github.com/Luxurioust/excelize"
	"log"
	"os"
	"strings"
)

func main() {

	nowFilePath, _ := os.Getwd()
	fileNames := strings.Split(nowFilePath, "/")
	fileName := fileNames[len(fileNames)-1]
	fmt.Println(fileNames)
	fmt.Println(nowFilePath)
	fmt.Println(fileName)

	fmt.Println("========================")
	fmt.Println("=生成目录文件=")
	fmt.Println("========================")
	fmt.Println("导出该目录下所有目录文件及子目录文件请输入: 1")
	fmt.Println("导出该目录下所有目录文件请输入: 2")

	var findType string
	// 用户输入信息
	findType = mytool.GetMsg()

	tree, _ := getDirTree.InitDirTree(nowFilePath, findType)
	//tree,err := getDirTree.InitDirTree(nowFilePath)
	//tree := getDirTree.Init()

	//fmt.Println(tree)
	//创建excel文件
	xlsx := excelize.NewFile()

	nameNum := 0
	nameEnd := "下所有文件.xlsx"
	var strNameNum string
	newName := nowFilePath + "/" + fileName + nameEnd
	_, err := os.Stat(newName)
	//fmt.Println(err)
	//fmt.Println(err == nil)
	//os.Exit(2)
	if err == nil {
		for err == nil {
			nameNum++
			strNameNum = fmt.Sprintf("%s(%d)%s", fileName, nameNum, nameEnd)
			_, err = os.Stat(strNameNum)

		}
	}

	if nameNum != 0 {
		fileName = fmt.Sprintf("%s(%d)", fileName, nameNum)
	}

	//fmt.Println(nameNum)
	//fmt.Println(fileName)
	//fmt.Println(strNameNum)
	//os.Exit(21)
	//创建新表单
	index := xlsx.NewSheet("目录整理")
	excelTree := make(map[string]string)
	for _, v := range tree {
		//fmt.Println(v)
		for im, iv := range v {
			//fmt.Println(im,iv)
			excelTree[im] = iv
			xlsx.SetCellValue("目录整理", im, iv)
		}
	}
	//设置默认打开的表单
	xlsx.SetActiveSheet(index)

	//保存文件到指定路径
	xlsxErr := xlsx.SaveAs("./" + fileName + "下所有文件.xlsx")
	if xlsxErr != nil {
		log.Fatal(xlsxErr)
	}
	//fmt.Println(excelTree)
}
