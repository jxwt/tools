package tools

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/astaxie/beego/logs"
	"io"
	"io/ioutil"
	"mime/multipart"
	"path"
	"strings"
)

// AliOssTool .
type AliOssTool struct {
	OssAliyunEndpoint        string
	OssAliyunAccessKeyId     string
	OssAliyunAccessKeySecret string
	OssAliyunBucket          string
}

// 校验文件上传类型
func GetAllowExt(fileExt string, fileTypes ...string) bool {
	allowExts := map[string]map[string]bool{
		"photo": {
			".jpg":  true,
			".jpeg": true,
			".png":  true,
		},
		"photos": {
			".jpg":  true,
			".jpeg": true,
			".png":  true,
		},
	}
	for _, fileType := range fileTypes {
		if allowExts[fileType] == nil {
			return true
		}
		if allowExts[fileType][fileExt] {
			return true
		}
	}
	return false
}

// 上传文件
// filePath 为空则上传到默认路径
func (c *AliOssTool) UploadFile(file multipart.File, fileName string, key string) (string, error) {
	defer file.Close()
	ext := path.Ext(fileName)
	if !GetAllowExt(ext, key) {
		return "", errors.New("上传文件类型错误，上传失败")
	}
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		logs.Warning(err)
		return "", errors.New("文件上传失败")
	}

	filePath := fmt.Sprintf("%x", GetRandomString(10)) + fileName
	return c.AliOssUpload(filePath, bytes.NewReader(buff))
}

func (c *AliOssTool) AliOssUpload(fileName string, reader io.Reader) (fileUrl string, err error) {
	// 创建OSSClient实例。
	client, err := oss.New(c.OssAliyunEndpoint, c.OssAliyunAccessKeyId, c.OssAliyunAccessKeySecret)
	if err != nil {
		logs.Warning(err)
		return "", err
	}
	bucket, err := client.Bucket(c.OssAliyunBucket)
	if err != nil {
		logs.Warning(err)
		return "", err
	}
	err = bucket.PutObject(fileName, reader)
	if err != nil {
		logs.Warning(err)
		return "", err
	}
	ossFilePath := c.OssAliyunBucket + "." + c.OssAliyunEndpoint + "/" + fileName
	if !strings.Contains(ossFilePath, "http://") {
		ossFilePath = "https://" + ossFilePath
	}
	return ossFilePath, nil
}

func (c *AliOssTool) AliOssDownload(fileName string, bucketName string) (io.Reader, error) {
	// 创建OSSClient实例。
	endpoint := "oss-cn-beijing.aliyuncs.com"
	if bucketName == "" {
		bucketName = c.OssAliyunBucket
		endpoint = c.OssAliyunEndpoint
	}
	client, err := oss.New(endpoint, c.OssAliyunAccessKeyId, c.OssAliyunAccessKeySecret)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	body, err := bucket.GetObject(fileName)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, body)
	return buf, nil
}
