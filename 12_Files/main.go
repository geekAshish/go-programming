package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("ex.txt")
	if err != nil {
		panic(err)
	}
	
	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	defer f.Close()

	fmt.Println(fileInfo.Name(), fileInfo.IsDir(), fileInfo.Size(), fileInfo.Mode().Perm(), fileInfo.ModTime())


	// Read File
	// 1. 
	buf := make([]byte, 10)

	bufLength, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buf); i++ {
		println("string data", bufLength, string(buf[i]))
	}

	fmt.Println("Data", bufLength, buf)


	// 2. we can read file like that but it's read file in one go in the memory, which will not good if the file is give
	fileData, err := os.ReadFile("ex.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(fileData, string(fileData))






	// Read Folder
	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	// fileInfoDir, err := dir.ReadDir(1)
	fileInfoDir, err := dir.ReadDir(-1)

	for _, fi := range fileInfoDir {
		fmt.Println(fi.Name(), fi.IsDir())
	}



	// create file
	file, err := os.Create("name.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// add content
	file.WriteString("Suii")
	file.WriteString("Cristiano ronaldo")



	// bytes writing
	bytes := []byte("messi")
	file.Write(bytes)



	// read and write to another file (streaming fashion)

	sourceFile, e := os.Open("ex.txt")
	desFile, ed := os.Open("name.txt")
	if e != nil || ed != nil {
		panic(e)
	}

	defer sourceFile.Close()
	defer desFile.Close()
	
	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(desFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}

			break
		}

		writeErr := writer.WriteByte(b)
		if writeErr != nil {
			panic(writeErr)
		}
	}

	writer.Flush()

	fmt.Println("written to new file successfully")



	// delete file
	removeErr := os.Remove("file_name.txt")
	if removeErr != nil {
		fmt.Println(removeErr)
	}
	fmt.Println("File removed successfully")



}