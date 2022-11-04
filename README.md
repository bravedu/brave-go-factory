# Brave-go-factory
是快速开发go项目的脚手架，集成go 语言gin框架。致力于让研发更加高效、结构更加清晰


## 项目打包 & 启动命令

```cgo
GOOS=linux GOARCH=amd64 go build -o 路径
```
### 环境公共变量
project_name=my_project
### 正式环境
app_env=online
### 仿真环境
app_env=beta
### 测试环境
app_env=test
### 开发环境
app_env=dev

```cgo
//按照不同环境变量替换即可
APP_ENV=$app_env PROJECT_NAME=${project_name} logs_path=$log_path nohup ./${project_name} 2> ${log_path}/api_log2.out 1> ${log_path}/api_log1.out &

```

#目录结构
```cgo
├── README.md
├── app
│   ├── controller #存项目控制器模块
│   │   ├── hello_controller.go
│   │   └── rpcresult.go
│   ├── models
│   ├── service #存项目服务层模块
│   │   └── hello_service.go
│   └── typespec #存放控制方法验证和返回描述模块
│       ├── base_spec.go
│       └── hello_spec.go
├── config #此模存放所有服务依赖的配置初始化
│   ├── config.go
│   ├── init_base.go
│   ├── init_jwt.go
│   ├── init_mysql.go
│   ├── init_nacos.go
│   ├── init_project_cnf.go  #项目基础配置
│   └── init_yaml.go         #yaml配置
├── constants #存放系统常量 和 错误常量配置
│   ├── constant.go
│   └── errors.go
├── dev_conf.yaml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── logs #项目日志存放位置
│   ├── project-app.log
│   ├── project-web.log
│   └── tracking-web.log
├── main.go
├── modules #项目模块画业务
│   ├── const_mod
│   │   ├── body_map.go
│   │   └── body_map_test.go
│   ├── jwt_mod
│   │   └── jwt.go
│   ├── log_mod
│   └── mysql_mod
│       ├── db_dao
│       │   ├── db_cnf.go
│       │   ├── db_interface.go
│       │   ├── db_store.go
│       │   └── goods.go
│       └── db_struct
│           ├── base.go
│           └── goods.go
├── pkg #存放项目使用工具化模块
│   ├── aes
│   │   ├── aes_cbc.go
│   │   ├── aes_ecb.go
│   │   ├── aes_gcm.go
│   │   ├── aes_test.go
│   │   └── pkcs_padding.go
│   ├── errgroup
│   │   ├── errgroup.go
│   │   └── errgroup_test.go
│   ├── logger
│   │   ├── logger.go
│   │   ├── logger_struct.go
│   │   └── recover.go
│   ├── util
│   │   ├── bind.go
│   │   ├── common.go
│   │   ├── convert.go
│   │   ├── convert_test.go
│   │   ├── files.go
│   │   ├── jwt.go
│   │   ├── random.go
│   │   └── string.go
│   ├── xhttp
│   │   ├── client.go
│   │   ├── client_test.go
│   │   └── model.go
│   ├── xlog
│   │   ├── color.go
│   │   ├── debug_logger.go
│   │   ├── error_logger.go
│   │   ├── info_logger.go
│   │   ├── log.go
│   │   ├── log_test.go
│   │   └── warn_logger.go
│   └── xtime
│       ├── constants.go
│       ├── data_time.go
│       ├── parse_format.go
│       ├── xtime.go
│       └── xtime_test.go
└── router #项目路由存放模块
    ├── api.go
    └── router.go

```


# 目录命名规范
## 控制器命名
### 1、文件命名 (以 _controller 为后缀，例如: 要定义一个名为 Hello 的控制器 )
```cgo
    hello_controller.go
```
### 2、控制方法命名(以 控制器名为前缀，例如: 要定义一个名为 Lists 的控制器方法)
```cgo
    func HelloList(c *gin.Context) {
        var (
            req  typespec.HelloListReq
            resp typespec.HelloListResp
        )
        if err := c.ShouldBind(&req); err != nil {
            c.JSON(http.StatusOK, WriteResponse(constants.ParamsValidateFail, err, nil))
            return
        }
        err := service.HelloListSvc(&req, &resp)
        if err != nil {
            c.JSON(http.StatusInternalServerError, WriteResponse(constants.ServerError, err))
            return
        }
        c.JSON(http.StatusOK, WriteResponse(constants.Success, nil, &resp))
    }
```

## 服务层文件命名
### 1、文件命名(以 _service 为后缀)
```cgo
    hello_service.go
```

### 2、服务方法创建以(控制器模块名 + 方法名 + Svc结尾, 例:要定义一个名为 Lists服务方法,为Hello控制器下 Lists 方法服务)
```cgo
    func HelloListSvc(req *typespec.HelloListReq, resp *typespec.HelloListResp) error {
        return nil
    }
```

## 控制器方法, request结构体 和 response定义

### 1、文件名定义(以 控制器方法 + _spec结尾)
```cgo
    hello_spec.go
```
### 2、request构体定义（以控制器名+控制方法名+Req 结尾,例如为Hello 控制器的 Lists 方法添加Req)
```cgo
    type HelloListReq struct {
        PageReq
        Id int `json:"id" form:"id" binding:"required"`
    }
```
### 3、response构体定义，(以控制器名+控制方法名+Resp 结尾,例如为Hello 控制器的 Lists 方法添加Resp)
```cgo
    type HelloListResp struct {
        List []HelloList `json:"list"`
    }
    
    type HelloList struct {
        Id    int    `json:"id"`
        Title string `json:"title"`
    }
```


