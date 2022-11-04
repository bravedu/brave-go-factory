package controller

import (
	"errors"
	"github.com/bravedu/brave-go-factory/constants"
	"github.com/gin-gonic/gin"
	"github.com/thinkeridea/go-extend/exnet"
	"strconv"
)

type Page struct {
	PageSize int `json:"page_size"`
	Offset   int `json:"offset"`
}

func NewPage(c *gin.Context) *Page {
	size, _ := strconv.ParseInt(c.Query("page_size"), 10, 64)
	offset, _ := strconv.ParseInt(c.Query("offset"), 10, 64)
	return &Page{
		PageSize: int(size),
		Offset:   int(offset),
	}
}

type Result struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func WriteResponse(code int, err error, data ...interface{}) *Result {
	if err == nil {
		err = errors.New(constants.CustomStatusText[code])
	} else {
		//自定义code
		errCode, _ := strconv.Atoi(err.Error())
		if errTxt, ok := constants.CustomStatusText[errCode]; errCode > 0 && ok {
			code = errCode
			return &Result{
				Code:  code,
				Error: errTxt,
				Data:  data,
			}
		}
	}
	if len(data) > 0 {
		return &Result{
			Code:  code,
			Error: err.Error(),
			Data:  data[0],
		}
	}
	return &Result{
		Code:  code,
		Error: err.Error(),
		Data:  data,
	}
}

func RequestIp(c *gin.Context) string {
	ip := exnet.ClientPublicIP(c.Request)
	if ip == "" {
		ip = exnet.ClientIP(c.Request)
	}
	return ip
}
