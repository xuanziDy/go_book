package handler

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xuanziDy/go_book/internal/domain"
	"github.com/xuanziDy/go_book/internal/service"
	_const "github.com/xuanziDy/go_book/internal/web/const"
	"github.com/xuanziDy/go_book/internal/web/request"
	"net/http"
)

type UserHandler struct {
	srv *service.UserService
}

func NewUserHandler(srv *service.UserService) *UserHandler {
	return &UserHandler{
		srv: srv,
	}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	//参数校验
	req, err := request.NewUserRequest(c).ValidateSignUp()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	//调用service进行注册逻辑
	err = h.srv.SignUp(c, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		//邮箱冲突要明确提示
		if err.Error() == _const.ErrUserDuplicateEmail {
			c.String(http.StatusOK, _const.ErrUserDuplicateEmail)
			return
		}
		c.String(http.StatusOK, _const.SystemError)
		return
	}

	c.String(http.StatusOK, _const.SinUpSuccess)
}

func (h *UserHandler) Login(c *gin.Context) {
	var err error

	//参数校验
	req, err := request.NewUserRequest(c).ValidateLogin()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	u, err := h.srv.Login(c, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	//账号密码不匹配
	if err != nil {
		if err.Error() == _const.ErrInvalidUserOrPassword {
			c.String(http.StatusOK, _const.ErrInvalidUserOrPassword)
			return
		}
		c.String(http.StatusOK, _const.SystemError)
		return
	}

	//存储session
	sess := sessions.Default(c)
	sess.Set("userId", u.Id)
	sess.Options(sessions.Options{
		Secure:   true,
		HttpOnly: true,
		MaxAge:   3600, //单位秒
	})
	sess.Save()

	c.String(http.StatusOK, _const.LoginSuccess)
}

func (h *UserHandler) Logout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Options(sessions.Options{
		MaxAge: -1,
	})
	sess.Save()
	c.String(http.StatusOK, _const.LogoutSuccess)
}

func (h *UserHandler) Edit(c *gin.Context) {
	var err error

	//参数校验
	req, err := request.NewUserRequest(c).ValidateEdit()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	userId, err := h.getUserId(c)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = h.srv.Edit(c, domain.User{
		Id:           userId,
		Nickname:     req.Nickname,
		Birthday:     req.Birthday,
		Introduction: req.Introduction,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, _const.EditFailed)
	}
	c.String(http.StatusInternalServerError, _const.EditSuccess)
}

func (h *UserHandler) Profile(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	u, err := h.srv.Profile(c, userId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("邮箱:%s 昵称:%s 生日:%s 个人简介：%s", u.Email, u.Nickname, u.Birthday, u.Introduction))
}

func (h *UserHandler) getUserId(c *gin.Context) (int64, error) {
	sess := sessions.Default(c)
	id := sess.Get("userId")
	if id == nil {
		return 0, errors.New(_const.SessionExpired)
	}
	userId, ok := id.(int64)
	if !ok {
		//类型错误
		return 0, errors.New("会话ID类型异常")
	}
	return userId, nil
}
