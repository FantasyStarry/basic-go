package middleware

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) IgnorePaths(path ...string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path...)
	return l
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	// 用go的形式编码解码
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		// 不需要登录校验
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}
		//if ctx.Request.URL.Path == "/users/login" ||
		//	ctx.Request.URL.Path == "/users/signup" {
		//	return
		//}
		sess := sessions.Default(ctx)
		id := sess.Get("userId")
		if id == nil {
			// 没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 先拿到上一次的更新时间 update_time
		updateTime := sess.Get("update_time")
		sess.Set("userId", id)
		now := time.Now()
		if updateTime == nil {
			// 说明还没有刷新过，第一次登录的第一个请求，还没有刷新过
			sess.Set("update_time", now)
			if err := sess.Save(); err != nil {
				panic(err)
			}
			// 设置要放在session中的值
			sess.Options(sessions.Options{
				MaxAge: 60 * 24 * 7,
			})
			return
		}
		// updateTime存在
		updateTimeVal, ok := updateTime.(time.Time)
		// 系统错误
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if now.Sub(updateTimeVal) > 24*time.Hour {
			sess.Set("update_time", now)
			if err := sess.Save(); err != nil {
				panic(err)
			}
			// 设置要放在session中的值
			sess.Options(sessions.Options{
				MaxAge: 60 * 24 * 7,
			})
		}
	}
}
