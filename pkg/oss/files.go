package oss

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	supportImageExtNames = []string{".jpg", ".jpeg", ".png", ".ico", ".svg", ".bmp", ".gif"}
	supportAudioExtNames = []string{".mp3", ".aac", ".opus", ".wav", ".flac", ".ape", ".alac"}
	supportExcelExtNames = []string{".xlsx"}
)

func Merge(filepath, watermark string) (string, error) {
	imgb, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer imgb.Close()

	img, _, err := image.Decode(imgb)
	if err != nil {
		return "", err
	}

	b := img.Bounds()
	imgW := b.Size().X
	imgH := b.Size().Y

	imgW = 300
	imgH = 300

	newImg := ImageResize(img, imgW, imgH)
	m := image.NewRGBA(newImg.Bounds())

	wmb, err := os.Open(watermark)
	if err != nil {
		return "", err
	}
	newwater, err := png.Decode(wmb)
	if err != nil {
		return "", err
	}
	defer wmb.Close()
	offset := image.Pt(50, 50)

	draw.Draw(m, newImg.Bounds().Add(image.Pt(0, 0)), newImg, image.ZP, draw.Src)
	draw.Draw(m, newwater.Bounds().Add(offset), newwater, image.ZP, draw.Over)

	p := strings.TrimSuffix(path.Base(filepath), path.Ext(filepath))
	imgw, err := os.Create(path.Join(path.Dir(watermark), p) + "_player" + path.Ext(filepath))
	if err != nil {
		return "", err
	}
	jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
	defer imgw.Close()
	return imgw.Name(), nil
}

// 图片大小调整
func ImageResize(src image.Image, w, h int) image.Image {
	return resize.Resize(uint(w), uint(h), src, resize.Lanczos3)
}

func GetImg(url, distdir string) (string, error) {
	p := strings.Split(url, "/")
	var name string
	if len(p) > 1 {
		name = p[len(p)-1]
	}
	out, err := os.Create(path.Join(distdir, name))
	if err != nil {
		return "", err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(out, bytes.NewReader(pix))
	if err != nil {
		return "", err
	}
	return out.Name(), nil
}

func ImageValid(file *multipart.FileHeader, maxUploadSize int64) (string, error) {
	extname := path.Ext(file.Filename)
	if file.Size > maxUploadSize {
		return "", errors.New(fmt.Sprintf("文件超过%vMb)", maxUploadSize/1024/1024))
	}
	for i := 0; i < len(supportImageExtNames); i++ {
		if supportImageExtNames[i] == strings.ToLower(extname) {
			return FileHash(file) + extname, nil
		}
	}
	return "", errors.New("不支持的文件格式")
}

func ExcelValid(file *multipart.FileHeader, maxUploadSize int64) (string, error) {
	extname := path.Ext(file.Filename)
	if file.Size > maxUploadSize {
		return "", errors.New(fmt.Sprintf("文件超过%vMb)", maxUploadSize/1024/1024))
	}
	for i := 0; i < len(supportExcelExtNames); i++ {
		if supportExcelExtNames[i] == strings.ToLower(extname) {
			return FileHash(file) + extname, nil
		}
	}
	return "", errors.New("不支持的文件格式")
}

func FileValid(file *multipart.FileHeader, maxUploadSize int64) (string, error) {
	extname := path.Ext(file.Filename)
	if file.Size > maxUploadSize {
		return "", errors.New(fmt.Sprintf("文件超过%vMb)", maxUploadSize/1024/1024))
	}
	for i := 0; i < len(supportAudioExtNames); i++ {
		if supportAudioExtNames[i] == strings.ToLower(extname) {
			return FileHash(file) + extname, nil
		}
	}
	return "", errors.New("不支持的文件格式")
}

func FileHash(file *multipart.FileHeader) string {
	src, err := file.Open()
	if err != nil {
		return ""
	}
	defer src.Close()

	myHash := sha256.New()
	_, err = io.Copy(myHash, src)
	if err != nil {
		return ""
	}
	hash := myHash.Sum(nil)
	sha256string := hex.EncodeToString(hash)
	return sha256string
}

//从远程下载封面合成
func MergeUposs(oss *Oss, remote, watermark string) (string, error) {
	src, err := GetImg(remote, path.Dir(watermark))
	if err != nil {
		return "", err
	}
	distfile, err := Merge(src, watermark)
	if err != nil {
		return "", err
	}
	err = oss.UploadOss(path.Base(distfile), distfile)
	if err != nil {
		return "", err
	}
	file := path.Base(distfile)
	return file, nil
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
