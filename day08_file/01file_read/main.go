package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//操作文件

func main() {

	// 只读方式打开当前目录下的 test.txt 文件
	file, err := os.Open("D:\\goproject\\go_project\\day08_file\\01file_read/test.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 关闭文件
	defer file.Close()

	// 使用Read方法读取数据
	//var tmp = make([]byte, 128)
	//n, err := file.Read(tmp)
	//if err == io.EOF {
	//	fmt.Println("文件读完了")
	//	return
	//}
	//if err != nil {
	//	fmt.Println("read file failed, err:", err)
	//	return
	//}
	//fmt.Printf("读取了%d字节数据\n", n)
	//fmt.Println(string(tmp[:n]))



	// 循环读取文件
	//var content []byte
	//var tmp = make([]byte, 128)
	//for {
	//	n, err := file.Read(tmp)
	//	if err == io.EOF {
	//		fmt.Println("文件读完了")
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("read file failed, err:", err)
	//		return
	//	}
	//	content = append(content, tmp[:n]...)
	//}
	//fmt.Println(string(content))

	//reader := bufio.NewReader(file)
	//for {
	//	line, err := reader.ReadString('\n') //注意是字符
	//	if err == io.EOF {
	//		if len(line) != 0 {
	//			fmt.Println(line)
	//		}
	//		fmt.Println("文件读完了")
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("read file failed, err:", err)
	//		return
	//	}
	//	fmt.Print(line)
	//}




	content, err := ioutil.ReadFile("D:\\goproject\\go_project\\day08_file\\01file_read/test.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))






}
