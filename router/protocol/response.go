/**
 * @Author: alessonhu
 * @Description:
 * @File:  response
 * @Version: 1.0.0
 * @Date: 2022/5/5 16:05
 */

package protocol

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// APIResponse 响应信息
type APIResponse struct {
	RetCode    int         `json:"ret"`
	ErrorCode  int         `json:"errorcode"`
	Msg        string      `json:"msg"`
	StackTrace string      `json:"stacktrace"`
	Data       interface{} `json:"data,omitempty"`
}

// MakeApiRsp 响应函数
func MakeApiRsp(c *gin.Context, errno *Errno, stacktrace string, data interface{}) {
	var rsp APIResponse
	rsp.ErrorCode = errno.Code
	rsp.Msg = errno.Message
	rsp.StackTrace = stacktrace
	rsp.Data = data
	c.JSON(http.StatusOK, rsp)
}
