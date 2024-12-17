package userError

/*
TODO: 这个包可以移除
	通过使用errors.New去管理特定错误的生成
*/

// 穷举所有可能的用户操作"错误"
const (
	paramsError   = "用户信息参数错误"
	accountError  = "用户名错误"
	passwordError = "用户密码错误"
	notLogin      = "用户未登录"
	banned        = "用户黑名单"
)

/*
userError
定义所有用户有关操作的错误信息
*/
type Error struct {
	msg string
}

func (userError *Error) Error() string {
	return userError.msg
}

// New 为什么不在这个指定"New"是UserError结构体的构造函数呢？
func AccountError() *Error {
	return &Error{ // 实例化一个新的UserError的结构体对象
		msg: accountError,
	}
}
func PasswordError() *Error {
	return &Error{ // 实例化一个新的UserError的结构体对象
		msg: passwordError,
	}
}
func ParamsError() *Error {
	return &Error{ // 实例化一个新的UserError的结构体对象
		msg: paramsError,
	}
}
func NotLogin() *Error {
	return &Error{ // 实例化一个新的UserError的结构体对象
		msg: notLogin,
	}
}
func Banned() *Error {
	return &Error{ // 实例化一个新的UserError的结构体对象
		msg: banned,
	}
}
