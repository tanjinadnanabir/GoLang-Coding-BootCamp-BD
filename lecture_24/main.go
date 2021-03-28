package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// dir, err := os.Getwd()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(dir)

	// createFile("master.txt", "Hello GoLang!")

	// isErr := createFile("master_1.txt", "Hello Bangladesh!")

	// fmt.Println(isErr)

	// fi, err := os.Stat("master_1.txt")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(fi.IsDir())
	// fmt.Println(fi.ModTime().Date())
	// fmt.Println(fi.Name())
	// fmt.Println(fi.Size())

	// how to make a folder

	// err := os.Mkdir("master", 0777)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// base := filepath.Base("./master")
	// fmt.Println(base)

	// base := filepath.Base(dir)
	// fullPath := filepath.Join("master")
	// fmt.Println(base)
	// fmt.Println(fullPath)
	
	// fullpath = E:\go\src\CodingBootCampGoLang\lecture_24

	// relativePath := filepath.Join("master")
	// absolutePath, _ := filepath.Abs("master")
	// newPath := filepath.Join(absolutePath, "..", "..", "lecture_25")

	// fmt.Println(base)
	// fmt.Println(relativePath)
	// fmt.Println(absolutePath)
	// fmt.Println(newPath)

	// os.Mkdir(newPath, 777)

	// os.Mkdir(`D:\TEST`, 777)
	// os.Rename(`D:\TEST`, `D:\TEST_01`)
	
	// *** End *** //


	// f := excelize.NewFile()

	// Create a new sheet.

	// index := f.NewSheet("Sheet2")

	// Set value of a cell.

	// f.SetCellValue("Sheet2", "A2", "Hello world.")
	// f.SetCellValue("Sheet1", "B2", 100)

	// Set active sheet of the workbook.
	// f.SetActiveSheet(index)

	// Save spreadsheet by the given path.
	// if err := f.SaveAs("test.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }

	// f, err := excelize.OpenFile("test.xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Get value from cell by given worksheet name and axis.

	// cell, err := f.GetCellValue("Sheet2", "A2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(cell)

	// Get all the rows in the Sheet1.
	// rows, err := f.GetRows("Sheet1")
	// for _, row := range rows {
	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }
}

func createFile(fileName, content string) bool {

	posf, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer posf.Close()

	_, err = posf.Write([]byte(content))

	// fmt.Println(n, err)
	if err != nil {
		return false
	}
	return true
}
