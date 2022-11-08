package controller

import (
	"github.com/bravedu/brave-go-factory/app/service"
	"github.com/bravedu/brave-go-factory/app/typespec"
	"github.com/bravedu/brave-go-factory/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HelloList
// @Summary 接口名称
// @schemes https
// @Description
// @Tags
// @Param activeId query int true "活动id"
// @Param page_size query int false "每页显示数量"
// @Param offset query int false "页码"
// @Accept json
// @Produce json
// @Success 200 {object} Result{data=typespec.HelloListResp} "成功"
// @Failure 10004 {object} Result
// @Router /v1/hello [get]
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

// HelloV2
// @Summary desc
// @schemes https
// @Description
// @Tags
// @name BraveDu
// @Accept json
// @Produce json
// @Param Content-Type header string false "参数类型" default("application/json")
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object body typespec.LoginReq true "手机号"
// @Security ApiKeyAuth
// @Success 200 {object} Result{data=typespec.LoginResp} "成功"
// @Failure 10004 {object} Result "失败, code状态码非0,显示error信息即可"
// @Router /v1/url [post]
func HelloV2(c *gin.Context) {
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
