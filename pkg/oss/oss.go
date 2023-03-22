package oss

import (
	"encoding/base64"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Oss struct {
	Bucket *oss.Bucket
}

func (o *Oss) Init(endpoint, accessKeyId, accessKeySecret, bucketName string) {
	// 创建OSSClient实例
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}

	// 获取存储空间
	o.Bucket, err = client.Bucket(bucketName)
	if err != nil {
		panic(err)
	}
}

func (o *Oss) UploadOss(objectName, localFileName string) error {
	// 上传文件。
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	err := o.Bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		return err
	}
	return nil
}

// 合并图片和水印，并上传到OSS
func (o *Oss) MergeUposs(cover, watermark string) (string, error) {
	distfile, err := Merge(cover, watermark)
	if err != nil {
		return "", err
	}
	err = o.UploadOss(path.Base(distfile), distfile)
	if err != nil {
		return "", err
	}
	defer os.Remove(distfile)
	fileName := path.Base(distfile)
	return fileName, nil
}

/*
@name 上传文件-字符串
*/
func (o *Oss) UploadByString(imageBase64 string, fileName string, fileType string) (url string, err error) {
	//获取文件名称
	if fileName == "" {
		if fileType == "" {
			fileType = "png"
		}
		fileName = fmt.Sprintf("%v", time.Now().UnixNano()) + "." + fileType //代码生成图片名称
	}
	filePath := fmt.Sprintf("%v/%v", time.Now().Format("2006-01-02"), fileName)
	uploadPath := filepath.Join("", filePath) //生成oos图片存储路径

	//获取图片内容并base64解密
	uploadString, _ := base64.StdEncoding.DecodeString(imageBase64)
	//创建OSSClient实例
	err = o.Bucket.PutObject(uploadPath, strings.NewReader(string(uploadString)))
	if err != nil {
		return url, err
	}
	return uploadPath, nil
}

/*
@name 上传文件-字符串
*/
func (o *Oss) UploadByImgByte(imageBytes []byte, fileName string, fileType string) (url string, err error) {
	//获取文件名称
	if fileName == "" {
		if fileType == "" {
			fileType = "png"
		}
		fileName = fmt.Sprintf("%v", time.Now().UnixNano()) + "." + fileType //代码生成图片名称
	}
	filePath := fmt.Sprintf("%v/%v", "user_qrcode", fileName)
	uploadPath := filepath.Join("", filePath) //生成oos图片存储路径
	//创建OSSClient实例
	err = o.Bucket.PutObject(uploadPath, strings.NewReader(string(imageBytes)))
	if err != nil {
		return url, err
	}
	return uploadPath, nil
}
