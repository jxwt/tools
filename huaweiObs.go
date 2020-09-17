package tools

import (
	"errors"
	"github.com/jxwt/tools/obs"
	"io"
	"strings"
)

var obsHandle Obs


type Obs struct {
	Bucket string
	Client  *obs.ObsClient
	Endpoint string
	Prefix string
}


func InitObs(accessKey, secretKey, endpoint, bucket string) {
	if obsHandle.Client == nil {
		c, err := obs.New(accessKey, secretKey, endpoint)
		if err != nil {
			panic(err)
		}
		obsHandle.Endpoint = endpoint
		obsHandle.Client = c
		obsHandle.Bucket = bucket
		obsHandle.Prefix = "https://" + obsHandle.Bucket + "." + obsHandle.Endpoint[strings.Index(obsHandle.Endpoint, "o"):] + "/"
	}
}

// obs上传文件
func ObsUploadFile(name string, reader io.Reader) (url string, err error) {
	input := &obs.PutObjectInput{}
	input.Bucket = obsHandle.Bucket
	input.Key = name
	input.Metadata = map[string]string{"meta": "value"}
	input.Body = reader
	_, err = obsHandle.Client.PutObject(input)
	if err != nil {
		return "", err
	}
	if obsError, ok := err.(obs.ObsError); ok {
		return "", errors.New(obsError.Message)
	}
	url = obsHandle.Prefix + name
	return
}

// obs下载文件
func ObsDownloadFile(name string) (io.ReadCloser, error) {
	input := &obs.GetObjectInput{}
	input.Bucket = obsHandle.Bucket
	input.Key = name
	output, err := obsHandle.Client.GetObject(input)
	if err != nil {
		return nil, err
	}
	return output.Body, nil
}