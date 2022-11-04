package controller

import (
	"github.com/bravedu/brave-go-factory/app/service"
	"github.com/bravedu/brave-go-factory/app/typespec"
	"github.com/bravedu/brave-go-factory/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HelloList doc
// @Summary 注释标题
// @Schemes https
// @Description
// @Accept json
// @Produce json
// @Param id  path  int  true  "参数ID"
// @Success 200 {object} Result{data=typespec.HelloListResp} "成功"
// @Router /v1/hello-lists [get]
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
// @Summary 测试post请求
// @schemes https
// @Description
// @Tags
// @name Authorization
// @Param activeId query int true "活动id"
// @Param page_size query int true "每页显示数量"
// @Param offset query int true "页码"
// @Accept json
// @Produce json
// @Success 200 {object} Result{data=typespec.HelloListResp} "成功"
// @Failure 10004 {object} Result
// @Router /v2/hello-v2 [post]
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
