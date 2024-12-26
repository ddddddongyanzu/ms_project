package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"test.com/project-api/pkg/model/user"
	common "test.com/project-common"
	"test.com/project-common/errs"
	login "test.com/project-grpc/user/login"
	"time"
)

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	// 初始化grpc客户端链接
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	rsp, err := LoginServiceClient.GetCaptcha(c, &login.CaptchaMessage{
		Mobile: mobile,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(rsp.Code))
}

func (u *HandlerUser) register(ctx *gin.Context) {
	// 1. 接收参数 参数模型
	result := &common.Result{}
	var req user.RegisterReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		//code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	// 2. 校验参数 判断参数是否合法
	if err := req.Verify(); err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
		return
	}
	// 3. 调用user grpc 服务 获取响应
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &login.RegisterMessage{}
	err = copier.Copy(msg, req)
	if err != nil {
		//code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy有误"))
		return
	}
	_, err = LoginServiceClient.Register(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	// 4. 返回结果
	ctx.JSON(http.StatusOK, result.Success(""))
}

func (u *HandlerUser) login(ctx *gin.Context) {
	// 1. 接收参数 参数模型
	result := &common.Result{}
	var req user.LoginReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		//code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	// 2. 调用user grpc 完成登录
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &login.LoginMessage{}
	err = copier.Copy(msg, req)
	if err != nil {
		//code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy有误"))
		return
	}
	loginRsp, err := LoginServiceClient.Login(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	// var rsp *user.LoginRsp
	rsp := &user.LoginRsp{}
	err = copier.Copy(&rsp, loginRsp)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy有误"))
		return
	}
	// 3. 返回结果
	ctx.JSON(http.StatusOK, result.Success(rsp))
}
