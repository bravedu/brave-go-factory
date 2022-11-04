package util

import "os"

type File struct {
	Name    string `json:"name"`
	Content []byte `json:"content"`
}

//
// GetEnvDefault
//  @-Description: 检查环境变量,如果不存在返回本身
//  @-param key
//  @-param defVal
//  @-return string
//
func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}

//
// CheckDirAndCreate
//  @-Description: 检查目录,如果不存在则创建
//  @-param path
//  @-return error
//
func CheckDirAndCreate(path string) error {
	if _, err := os.Stat(path); err == nil {
		//存在
		return nil
	} else {
		//不存在
		err := os.MkdirAll(path, 0751)
		if err != nil {
			return err
		}
		return nil
	}
}
