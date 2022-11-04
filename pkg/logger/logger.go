package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bravedu/brave-go-factory/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var AppLog *logrus.Logger
var WebLog *logrus.Logger
var TrackingLog *logrus.Logger

func init() {
	initAppLog()
	initWebLog()
	initTrackingLog()
}

/**
初始化AppLog
*/
func initAppLog() {
	logFileName := "project-app.log"
	AppLog = initLog(logFileName)
}

/**
初始化WebLog
*/
func initWebLog() {
	logFileName := "project-web.log"
	WebLog = initLog(logFileName)
}

/**
初始化TrackingLog埋点日志
*/
func initTrackingLog() {
	logFileName := "tracking-web.log"
	TrackingLog = initLog(logFileName)
}

/**
初始化日志句柄
*/
func initLog(logFileName string) *logrus.Logger {
	logPath := util.GetEnvDefault("logs_path", "logs/")
	util.CheckDirAndCreate(logPath)
	// 日志文件
	logName := path.Join(logPath, logFileName)
	var f *os.File
	var err error
	//判断日志文件是否存在，不存在则创建，否则就直接打开
	if _, err := os.Stat(logName); os.IsNotExist(err) {
		f, err = os.Create(logName)
	} else {
		f, err = os.OpenFile(logName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}
	if err != nil {
		fmt.Println("open log file failed")
	}

	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	}
	log.Out = f
	//// 设置日志级别
	log.SetLevel(logrus.InfoLevel)
	// 设置 rotatelogs
	/*
		logWriter, _ := rotatelogs.New(
			// 分割后的文件名称
			logName+".%Y%m%d.log",

			// 生成软链，指向最新日志文件
			rotatelogs.WithLinkName(logName),

			// 设置最大保存时间(7天)
			rotatelogs.WithMaxAge(7*24*time.Hour),

			// 设置日志切割时间间隔(1天)
			rotatelogs.WithRotationTime(24*time.Hour),
		)

		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}
		log.SetReportCaller(true)
		lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})

		// 新增钩子
		log.AddHook(lfHook)
	*/
	return log
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

/**
Gin中间件函数，记录请求日志
*/
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}
		traceId := fmt.Sprintf("%s-%v", util.GenIdCreate(20), time.Now().UnixMilli())
		WebLog.AddHook(NewTraceIdHook(traceId))
		AppLog.AddHook(NewTraceIdHook(traceId))
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		// 开始时间
		startTime := time.Now().UnixNano()
		postJsonData := ""

		if c.Request.Method == "POST" && c.Request.Header.Get("Content-Type") != "" && strings.ToLower(c.Request.Header.Get("Content-Type")) == "application/json" {
			body, err := c.GetRawData()
			if err != nil {
				fmt.Println(err.Error())
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			postJsonData = string(body)
		}

		// 处理请求
		c.Next()

		responseBody := bodyLogWriter.body.String()
		var responseCode string
		var responseMsg string
		//var responseData interface{}

		if responseBody != "" {
			res := Result{}
			err := json.Unmarshal([]byte(responseBody), &res)
			if err == nil {
				responseCode = strconv.Itoa(res.Code)
				responseMsg = res.Error
				//responseData = res.Data
			}
		}
		// 结束时间
		endTime := time.Now().UnixNano()
		if c.Request.Method == "POST" && postJsonData == "" {
			c.Request.ParseForm()
			postJsonData = c.Request.PostForm.Encode()
			if postJsonData == "" {
				postJsonData = "got nothing form data, pls check params type and contentType"
			}
		}

		costTime, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", float64(endTime-startTime)/1e6), 64)
		// 日志格式
		WebLog.WithFields(logrus.Fields{
			"request_method":    c.Request.Method,
			"request_uri":       c.Request.RequestURI,
			"request_proto":     c.Request.Proto,
			"request_useragent": c.Request.UserAgent(),
			"request_referer":   c.Request.Referer(),
			"request_post_data": postJsonData,
			"request_auth_data": c.GetInt("uid"),
			"request_client_ip": c.ClientIP(),

			"response_status_code": c.Writer.Status(),
			"response_code":        responseCode,
			"response_msg":         responseMsg,
			//"response_data":        responseData,
			"cost_time": costTime,
		}).Info("access")
	}
}

type TraceIdHook struct {
	TraceId string
}

func NewTraceIdHook(traceId string) logrus.Hook {
	hook := TraceIdHook{
		TraceId: traceId,
	}
	return &hook
}

func (hook *TraceIdHook) Fire(entry *logrus.Entry) error {
	entry.Data["Trace-Id"] = hook.TraceId
	return nil
}

func (hook *TraceIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
