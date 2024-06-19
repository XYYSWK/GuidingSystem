package middlewares

import (
	"GuidingSystem/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
)

/*
自定义的 recovery 异常捕获中间件
*/

// Recovery 捕获项目可能出现的 panic，并向开发者发送异常信息的邮件
func Recovery(stack bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 检查连接是否断开，因为这并不是真正需要进行恐慌堆栈跟踪的情况
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connections reset by peer") {
							brokenPipe = true
						}
					}
				}

				// 将请求对象转换为字节切片
				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				var body string
				data, ok := ctx.Get(Body)
				if ok {
					body = string(data.([]byte))
				}

				if brokenPipe {
					global.Logger.Error(ctx.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)))
					// 如果连接已断开，我们就无法写入状态
					ctx.Error(err.(error)) // 将错误信息与上下文关联
					ctx.Abort()            // 阻止调用后续的处理函数
					return
				}
				if stack { // 如果需要记录堆栈信息
					global.Logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack()))) // 记录当前 goroutine 的堆栈跟踪信息到日志中
				} else { // 不需要记录到堆栈信息
					global.Logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("body", body))
				}
				ctx.AbortWithStatus(http.StatusInternalServerError) //阻止调用后续的处理函数，并返回“服务器内部错误”的状态码
			}
		}()
		ctx.Next()
	}
}
