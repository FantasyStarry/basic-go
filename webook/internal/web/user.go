package web

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/service"
	"errors"
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// UserHandler 定义跟User用户有关的路由
type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		// 和上面比起来，用 ` 看起来就比较清爽
		passwordRegexPattern = `^(?=.*[!@#$%^&*()_+])(?=.*[a-zA-Z])(?=.*\d)[a-zA-Z\d!@#$%^&*()_+]{9,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

func (u *UserHandler) RegisterUsersRoutes(server *gin.Engine) {
	// 分组注册
	ug := server.Group("/users")
	//注册
	ug.POST("/signup", u.SignUp)
	//登录
	ug.POST("/login", u.Login)
	//编辑
	ug.POST("/edit", u.Edit)
	//个人信息
	ug.GET("/profile", u.Profile)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req SignUpReq
	// Bind 方法根据Content-Type来解析到你的数据req里面
	// 解析错了，就会直接返回一个400的错误
	if err := ctx.Bind(&req); err != nil {
		return
	}
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "你的邮箱格式不正确")
		return
	}
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次输入不密码不相同")
		return
	}
	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		//记录日志
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字，特殊字符")
		return
	}
	//调用一下service层
	err = u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if errors.Is(err, service.ErrUserDuplicateEmail) {
		ctx.String(http.StatusOK, "邮箱已注册")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}
	ctx.String(http.StatusOK, "注册成功")
}

func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	uRepo, err := u.svc.Login(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		ctx.String(http.StatusOK, "用户名或密码错误")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if uRepo == (domain.User{}) {
		ctx.String(http.StatusOK, "登录失败")
		return
	}
	// 登录成功之后
	sess := sessions.Default(ctx)
	// 设置要放在session中的值
	sess.Set("userId", uRepo.Id)
	sess.Options(sessions.Options{
		MaxAge: 60 * 24 * 7,
	})
	err = sess.Save()
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, "登录成功")
	return
}

func (u *UserHandler) Logout(ctx *gin.Context) {
	// 退出登录
	sess := sessions.Default(ctx)
	// 设置要放在session中的值
	sess.Options(sessions.Options{
		MaxAge: -1,
	})
	err := sess.Save()
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, "退出登录成功")
	return
}

func (u *UserHandler) Edit(ctx *gin.Context) {
	type EditReq struct {
		NickName string `json:"nickname"`
		Birthday string `json:"birthday"`
		AboutMe  string `json:"aboutMe"`
	}
	sess := sessions.Default(ctx)
	userId := sess.Get("userId").(int64)
	var req EditReq
	if err := ctx.Bind(&req); err != nil {
		ctx.String(http.StatusOK, "绑定数据失败")
		return
	}
	formatTime := "2006-01-02"
	birthday, err := time.Parse(formatTime, req.Birthday)
	if err != nil {
		ctx.String(http.StatusOK, "时间格式转化错误")
		fmt.Printf("%s", err.Error())
		return
	}
	err = u.svc.EditUserInfo(ctx, domain.User{
		Id:       userId,
		Nickname: req.NickName,
		Birthday: birthday,
		AutoMe:   req.AboutMe,
	})
	if err != nil {
		ctx.String(http.StatusOK, "修改个人信息失败")
		return
	}
	ctx.String(http.StatusOK, "修改个人信息成功")
}

func (u *UserHandler) Profile(ctx *gin.Context) {
	// 取出存放在session中的userId
	sess := sessions.Default(ctx)
	userId := sess.Get("userId").(int64)
	user, err := u.svc.Profile(ctx, userId)
	if err != nil {
		ctx.String(http.StatusOK, "个人信息获取失败")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Id":       user.Id,
		"Nickname": user.Nickname,
		"Email":    user.Email,
		"Birthday": user.Birthday,
		"AutoMe":   user.AutoMe,
	})
}
