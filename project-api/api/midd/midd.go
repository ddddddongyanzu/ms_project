package midd

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"test.com/project-api/api/user"
	common "test.com/project-common"
	"test.com/project-common/errs"
	login "test.com/project-grpc/user/login"
	"time"
)

func TokenVerify() func(*gin.Context) {
	return func(c *gin.Context) {
		result := common.Result{}
		// 1. 从header中获取token
		token := c.GetHeader("Authorization")
		// 2. 调用user服务及进行token认证
		ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancelFunc()
		response, err := user.LoginServiceClient.TokenVerify(ctx, &login.LoginMessage{Token: token})
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusOK, result.Fail(code, msg))
			c.Abort()
			return
		}
		// 3. 处理结果，认证通过 将信息放入gin的上下文 失败返回未登录
		c.Set("memberId", response.Member.Id)
		c.Next()
	}
}
