/**
 * @Author: alessonhu
 * @Description:
 * @File:  code
 * @Version: 1.0.0
 * @Date: 2022/5/5 16:08
 */

package protocol

var (
	OK                = &Errno{Code: 0, Message: "OK"}
	ErrRequest        = &Errno{Code: 10001, Message: "请求参数不合法"}
	ErrInternalServer = &Errno{Code: 10002, Message: "内部服务器错误"}
	ErrValidation     = &Errno{Code: 10003, Message: "签名校验失败"}
)

type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}
