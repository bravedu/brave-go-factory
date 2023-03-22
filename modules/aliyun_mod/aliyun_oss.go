package aliyun_mod

import (
	"github.com/bravedu/brave-go-factory/pkg/oss"
	"os"
	"sync"
)

var (
	ossPool *OssPool
	ossOnce sync.Once
)

type OssCnf struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	BucketName      string `yaml:"bucket_name"`
	AccessDomain    string `yaml:"access_domain"`
}

type OssUploadCnf struct {
	ImgDir            string `yaml:"img_dir"`
	FileDir           string `yaml:"file_dir"`
	WaterFile         string `yaml:"water_file"`
	ThumbUrlSuffix    string `yaml:"thumb_url_suffix"`
	ImgUrlPrefix      string `yaml:"img_url_prefix"`
	AudioUrlPrefix    string `yaml:"audio_url_prefix"`
	MaxFileUploadSize int64  `yaml:"max_file_upload_size"`
	MaxImgUploadSize  int64  `yaml:"max_img_upload_size"`
}

type OssPool struct {
	Cli *oss.Oss
}

func OssClient(cnf *OssCnf, uploadCnf *OssUploadCnf) *OssPool {
	ossOnce.Do(func() {
		ossClient := new(oss.Oss)
		ossClient.Init(cnf.Endpoint, cnf.AccessKeyId, cnf.AccessKeySecret, cnf.BucketName)
		ossPool = &OssPool{
			Cli: ossClient,
		}
		initUpload(uploadCnf)
	})
	return ossPool
}

func initUpload(cnf *OssUploadCnf) {
	if !oss.Exists(cnf.ImgDir) {
		err := os.MkdirAll(cnf.ImgDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	if !oss.Exists(cnf.FileDir) {
		err := os.MkdirAll(cnf.FileDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	//if !util.Exists(c.YamlDao.Upload.WaterFile) {
	//	panic(errors.New("找不到水印文件：" + c.YamlDao.Upload.WaterFile))
	//}
}
