package tools

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func DeCompress(tarFile, dest string) error {
	if strings.HasSuffix(tarFile, ".zip") {
		return zipDeCompress(tarFile, dest)
	}
	return nil
}

func zipDeCompress(zipFile, dest string) error {
	or, err := zip.OpenReader(zipFile)
	defer or.Close()
	if err != nil {
		return err
	}
	for _, item := range or.File {
		if item.FileInfo().IsDir() {
			os.Mkdir(dest+item.Name, 0777)
			continue
		}
		rc, _ := item.Open()
		dst, _ := createFile(dest + item.Name)
		payload, err := ioutil.ReadAll(rc)
		_, err = dst.Write(payload)
		if err != nil {
			log.Print(dest + item.Name)
			log.Print(err)
			return err
		}
	}
	return nil
}

func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}

// Zip 打包成zip文件
func Zip(srcDir string, zipFileName string) {
	// 预防：旧文件无法覆盖
	os.RemoveAll(zipFileName)
	// 创建：zip文件
	zipfile, _ := os.Create(zipFileName)
	defer zipfile.Close()
	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()
	// 遍历路径信息
	filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {
		// 如果是源路径，提前进行下一个遍历
		if path == srcDir {
			return nil
		}
		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, srcDir+`\`)
		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}
		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})
}