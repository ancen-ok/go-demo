package controller

import (
	"gitee.com/molonglove/goboot/gin"
	"go-demo/app/models/entity"
	"go-demo/app/models/request"
	"go-demo/app/models/response"
	"go-demo/app/models/vo"
	"go-demo/app/service"
	"net/http"
)

var User = new(UserController)

type UserController struct {
	BaseController
}

// CaptchaImage 获取验证码
func (u *UserController) CaptchaImage(c *gin.Context) {
	var (
		result *response.CaptchaImageResponse
		err    *response.BusinessError
	)
	if result, err = service.User.CaptchaImage(); err != nil {
		c.JSON(http.StatusOK, response.ResultCustom(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(result))
}

// Login 登陆
func (u *UserController) Login(c *gin.Context) {

	var (
		err       error
		param     request.UserLoginParam
		customErr *response.BusinessError
		result    *response.UserLoginResponse
	)
	if err = c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}

	if result, customErr = service.User.UserLogin(&param, c); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}

	c.JSON(http.StatusOK, response.OkMsg(response.LoginSuccess, result))
}

// GetUserInfo 获取用户登录信息
func (u *UserController) GetUserInfo(c *gin.Context) {
	var (
		userId    int64
		customErr *response.BusinessError
		result    *response.UserInfoResponse
		err       error
	)
	if userId, err = c.QueryInt64("userId"); err != nil {
		c.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}

	if result, customErr = service.User.GetUserInfo(userId); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	c.JSON(http.StatusOK, response.Ok(result))
}

// Page 分页
func (u *UserController) Page(c *gin.Context) {
	var (
		param     request.PageParam
		customErr *response.BusinessError
		result    *response.PageData
		err       error
	)
	if err = c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}

	if result, customErr = service.User.Page(&param); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
	}
	c.JSON(http.StatusOK, response.Ok(result))
}

// 用户创建
func (u *UserController) Create(c *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.UserCreateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = u.Parse(c, "用户创建", vo.Add, nil)
}
