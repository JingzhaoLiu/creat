package copy

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 复制文件夹
func CopyDir(src string, dest string) {
	srcOriginal := src
	err := filepath.Walk(src, func(src string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			//			fmt.Println(f.Name())
			//			copyDir(f.Name(), dest+"/"+f.Name())
		} else {
			// fmt.Println(src)
			// fmt.Println(srcOriginal)
			// fmt.Println(dest)
 
			destNew := strings.Replace(src, srcOriginal, dest, -1)
			// fmt.Println(destNew)
			fmt.Println("CopyFile:" + src + " to " + destNew)
			copyFile(src, destNew)
		}
		//println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
 

func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func copyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()
	fmt.Println("dst:" + dst)
	dstSlices := strings.Split(dst, "/")
	dstSlicesLen := len(dstSlices)
	destDir := ""
	for i := 0; i < dstSlicesLen-1; i++ {
		destDir = destDir + dstSlices[i] + "/"
	}
	//destDir := getParentDirectory(dst)
	fmt.Println("destDir:" + destDir)
	b, err := pathExists(destDir)
	if b == false {
		err := os.MkdirAll(destDir, os.ModePerm) //在当前目录下生成md目录
		if err != nil {
			fmt.Println(err)
		}
	}
	dstFile, err := os.Create(dst)
 
	if err != nil {
		fmt.Println(err.Error())
		return
	}
 
	defer dstFile.Close()
 
	return io.Copy(dstFile, srcFile)
}