package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode            int64 = 0
	ServiceErrCode         int64 = 10001
	ParamErrCode           int64 = 10002
	LoginErrCode           int64 = 10003
	UserNotExistErrCode    int64 = 10004
	UserNameHasUsedErrCode int64 = 10005
	TokenInvalidErrCode    int64 = 10006
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success            	= NewErrNo(SuccessCode, "成功")
	ServiceErr			= NewErrNo(ServiceErrCode, "服务启动失败")
	ParamErr			= NewErrNo(ParamErrCode, "参数有误")
	ErrService         error = errors.New("服务异常，请稍后再试")
	Errparameter       error = errors.New("参数不正确")
	ErrLogin           error = errors.New("用户名或密码错误")
	ErrUserNotExist    error = errors.New("用户不存在")
	ErrUserNameHasUsed error = errors.New("用户名已经被使用")
	ErrTokenInvalid    error = errors.New("token已过期或不可用，请重新登录")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
