package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	_const "github.com/xuanziDy/go_book/internal/web/const"
	"regexp"
	"unicode/utf8"
)

type UserRequest struct {
	Context *gin.Context
}

func NewUserRequest(context *gin.Context) *UserRequest {
	return &UserRequest{
		Context: context,
	}
}

type UserSignUpReq struct {
	Email           string `json:"email"`
	ConfirmPassword string `json:"confirmPassword"`
	Password        string `json:"password"`
}

func (r *UserRequest) ValidateSignUp() (UserSignUpReq, error) {

	//绑定参数到结构体
	var req UserSignUpReq
	if err := r.Context.Bind(&req); err != nil {
		return req, err
	}

	//验证邮箱
	ok := regexp.MustCompile(_const.EmailRegexPattern).MatchString(req.Email)
	if !ok {
		return req, errors.New(_const.ErrUserEmailFormat)
	}

	//验证密码
	if req.ConfirmPassword != req.Password {
		return req, errors.New(_const.ErrUserPasswordNotSame)
	}

	ok = regexp.MustCompile(_const.PasswordRegexPattern).MatchString(req.Password)
	if !ok {
		return req, errors.New(_const.ErrUserPasswordRule)
	}
	return req, nil
}

type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *UserRequest) ValidateLogin() (UserLoginReq, error) {
	var req UserLoginReq
	if err := r.Context.Bind(&req); err != nil {
		return req, err
	}

	ok := regexp.MustCompile(_const.EmailRegexPattern).MatchString(req.Email)
	if !ok {
		return req, errors.New(_const.ErrUserEmailFormat)
	}

	ok = regexp.MustCompile(_const.PasswordRegexPattern).MatchString(req.Password)
	if !ok {
		return req, errors.New(_const.ErrUserPasswordRule)
	}
	return req, nil
}

type UserEditReq struct {
	Nickname     string `json:"nickname"`
	Birthday     string `json:"birthday"`
	Introduction string `json:"introduction"`
}

func (r *UserRequest) ValidateEdit() (UserEditReq, error) {
	var err error

	var req UserEditReq
	if err = r.Context.Bind(&req); err != nil {
		return req, err
	}

	if req.Nickname != "" {
		if l := utf8.RuneCountInString(req.Nickname); l >= 10 {
			return req, errors.New(_const.ErrNickNameRule)
		}
	}
	if req.Birthday != "" {
		if ok := regexp.MustCompile(_const.BirthdayRegexPatten).MatchString(req.Birthday); !ok {
			return req, errors.New(_const.ErrBirthdayRule)
		}
	}

	if req.Introduction != "" {
		if l := utf8.RuneCountInString(req.Introduction); l >= 100 {
			return req, errors.New(_const.ErrIntroductionRule)
		}
	}

	return req, nil
}
