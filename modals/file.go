package modals

import (
	"fmt"
	"os"
)

// deleteFile func
func deleteFile() {
	userFile := "rayliao.txt"
	// 删除文件
	err := os.Remove(userFile)
	if err != nil {
		fmt.Println(err)
	}
}

// FileHandle func
func FileHandle() {
	os.Mkdir("rayliao", 0777)
	os.MkdirAll("rayliao/test1/test2", 0777)
	// 目录下有文件或其他目录会出错
	err := os.Remove("rayliao")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("rayliao")

	// 创建一个文件并写入
	userFile := "rayliao.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}

	// 读取文件
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		// 读取数据到buf
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}

	deleteFile()
}
