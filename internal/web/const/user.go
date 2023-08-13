package _const

// 文案
const (
	ErrUserDuplicateEmail    = "邮箱冲突"
	ErrUserEmailFormat       = "你的邮箱格式不对"
	ErrUserPasswordNotSame   = "两次输入的密码不一致"
	ErrUserPasswordRule      = "密码必须大于8位，包含数字、特殊字符"
	ErrInvalidUserOrPassword = "用户名或密码不对"

	ErrNickNameRule     = "昵称限制在10个字符内"
	ErrIntroductionRule = "个人简介限制在10个字符内"
	ErrBirthdayRule     = "生日需满足1996-01-02格式"

	LoginSuccess = "登录成功"
	LoginFailed  = "登录失败"
	SinUpFailed  = "注册失败"
	SinUpSuccess = "注册成功"

	SystemError   = "系统异常"
	LogoutSuccess = "退出登录成功"

	EditSuccess = "编辑成功"
	EditFailed  = "编辑失败"

	SessionExpired = "会话过期，请重新登录"
)

// 验证规则
const (
	EmailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	PasswordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	BirthdayRegexPatten  = "^d{4}-d{2}-d{2}$"
)
