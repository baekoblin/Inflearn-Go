package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 파일 읽기
	readFile()

	// 파일 쓰기
	writeFile()

	// 파일에 데이터 추가
	appendToFile()

	// 파일 정보 확인
	fileInfo()

	// 파일 존재 여부 확인
	checkFileExistence()

	// 디렉토리 생성
	createDirectories()

	// 디렉토리 내용 읽기
	readDirectory()

	// 디렉토리 삭제
	deleteFileAndDirectory()

	// 디렉토리 이동 및 이름 변경
	renameFileAndDirectory()
}

func readFile() {
	fmt.Println("Reading File")
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:", string(content))
}

func writeFile() {
	fmt.Println("Writing File")
	content := []byte("Hello, World!")
	err := ioutil.WriteFile("example.txt", content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File written successfully")
}

func appendToFile() {
	fmt.Println("Appending to File")
	file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("\nAppended text")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Text appended successfully")
}

func fileInfo() {
	fmt.Println("Getting File Info")
	fileInfo, err := os.Stat("example.txt")
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("File Size:", fileInfo.Size())
	fmt.Println("File Mode:", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
}

func checkFileExistence() {
	fmt.Println("Checking File Existence")
	_, err := os.Stat("example.txt")
	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
	} else {
		fmt.Println("File exists")
	}
}

func createDirectories() {
	fmt.Println("Creating Directories")
	err := os.Mkdir("example_dir", 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	err = os.MkdirAll("parent_dir/child_dir", 0755)
	if err != nil {
		fmt.Println("Error creating nested directories:", err)
		return
	}

	fmt.Println("Directories created successfully")
}

func readDirectory() {
	fmt.Println("Reading Directory")
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func deleteFileAndDirectory() {
	fmt.Println("Deleting File and Directory")
	err := os.Remove("example.txt")
	if err != nil {
		fmt.Println("Error removing file:", err)
		return
	}

	err = os.RemoveAll("example_dir")
	if err != nil {
		fmt.Println("Error removing directory:", err)
		return
	}

	fmt.Println("File and directory removed successfully")
}

func renameFileAndDirectory() {
	fmt.Println("Renaming File and Directory")
	err := os.Rename("example.txt", "new_example.txt")
	if err != nil {
		fmt.Println("Error renaming file:", err)
		return
	}

	err = os.Rename("example_dir", "new_example_dir")
	if err != nil {
		fmt.Println("Error renaming directory:", err)
		return
	}

	fmt.Println("File and directory renamed successfully")
}
