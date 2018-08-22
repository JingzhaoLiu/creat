package main

import (
	"os"
	"fmt"
	"copy/copy"
	"copy/decompress"
)

// 生成文件夹
func mkdir(name string) {
	err := os.Mkdir(name, 0777)
	
	if err != nil {
		fmt.Printf("%s 创建失败\n",name)
	}else{
		fmt.Printf("%s 创建成功\n",name)
	}
}
func main(){
	//  dirname <- os.Args[1]
	//  path <- os.Args[2]
	toPath := os.Args[2] + os.Args[1]
	
	// 创建总文件夹
	mkdir(toPath)
	// mkdir(toPath+"/bin")
	// copy.CopyDir("../toolkit/bin", toPath+"/bin")
	// mkdir(toPath+"/src")
	// mkdir(toPath+"/pkg")
	// copy.CopyDir("../toolkit/pkg", toPath+"/pkg")
	// mkdir(toPath+"/src/dirname")

	// 解压工具包
    path := "../toolkit"
	if _,err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			zipFile := "../toolkit.zip"
			dest := "../"
			decompress.DeCompress(zipFile, dest)
		} else {
			// other error
			fmt.Println("other error")
		}
	} else{
		fmt.Println("不需要解压")
	}
	
    // 拷贝工具包
	copy.CopyDir("../toolkit", toPath)
	
}



	
	
