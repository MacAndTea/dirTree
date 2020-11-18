package main

import (
	"dirTree/getDirTree"
	"fmt"
	"github.com/Luxurioust/excelize"
	"log"
	"os"
	"dirTree/mytool"
)

func main() {

	nowFilePath, _ := os.Getwd()

	fmt.Println("========================")
	fmt.Println("=生成目录文件=")
	fmt.Println("========================")
	fmt.Println("导出该目录下所有目录文件及子目录文件请输入: 1")
	fmt.Println("导出该目录下所有目录文件请输入: 2")

	var findType string
	// 用户输入信息
	findType = mytool.GetMsg()

	tree,_ := getDirTree.InitDirTree(nowFilePath,findType)
	//tree,err := getDirTree.InitDirTree(nowFilePath)
	//tree := getDirTree.Init()

	fmt.Println(tree)
	//创建excel文件
	xlsx := excelize.NewFile()

	//创建新表单
	index := xlsx.NewSheet("成绩表")
	excelTree := make(map[string]string)
	for _,v := range tree{
		//fmt.Println(v)
		for im,iv := range v{
			//fmt.Println(im,iv)
			excelTree[im] = iv
			xlsx.SetCellValue("成绩表", im, iv)
		}
	}
	//设置默认打开的表单
	xlsx.SetActiveSheet(index)

	//保存文件到指定路径
	err := xlsx.SaveAs("./目录下所有文件.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(excelTree)
}
