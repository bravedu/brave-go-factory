package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//定义 文件名、行号、方法名
				fileName, line, functionName := "?", 0, "?"
				pc, fileName, line, ok := runtime.Caller(5)
				if ok {
					functionName = runtime.FuncForPC(pc).Name()
					functionName = filepath.Ext(functionName)
					functionName = strings.TrimPrefix(functionName, ".")
				}
				AppLog.WithFields(logrus.Fields{
					"file": fmt.Sprintf("%s:%d", fileName, line),
					"func": functionName,
					"txt":  fmt.Sprintf("%s", r),
				}).Error("业务异常-最终拦截")
				c.JSON(http.StatusOK, &Result{
					Code:  recoverErrCode,
					Data:  nil,
					Error: "服务业务异常请联系技术伙伴",
				})
				return
			}
		}()
		c.Next()
	}
}
