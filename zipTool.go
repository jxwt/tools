package tools

import (
	"archive/zip"
	"io/ioutil"
	"log"
	"os"
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
