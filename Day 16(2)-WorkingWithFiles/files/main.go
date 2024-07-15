package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//Write in the File

	content := "The content to be written"

	file, err := os.Create("./goWriteFile.txt") // os package has a Create method to create a file at specifies location

	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	//To Write
	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length is : ", length)

	defer file.Close()
	readFile("./goWriteFile.txt") //We need to give full path of the File
}

func readFile(fileName string) {
	dataInBytes, err := ioutil.ReadFile(fileName) // Here at first we get the data in format of bytes and we need to convert it to string

	checkNilError(err)

	fmt.Println("Text Data inside the File :", dataInBytes)         // Here we get the data in bytes
	fmt.Println("Text Data inside the File :", string(dataInBytes)) //convert it into the string
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
