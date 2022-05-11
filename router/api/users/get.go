/**
 * @Author: alessonhu
 * @Description:
 * @File:  get
 * @Version: 1.0.0
 * @Date: 2022/5/5 17:26
 */

package users

import (
	"apidemo/logger"
	"apidemo/model"
	ctrl "apidemo/router/protocol"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetUserById doc
// @Summary 获取单个用户
// @Description 根据id获取指定用户
// @Tags
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Success 200 {string} string "{"Code":0,"Message":"OK","Data":{}}"
// @Failure 200 {string} string "{"Code":10001,"Message":"请求参数不合法"}"
// @Router /users/{id} [get]
func GetUserById(c *gin.Context) {
	retCode := ctrl.OK
	stacktrace := ""
	var data interface{}
	defer func() {
		ctrl.MakeApiRsp(c, retCode, stacktrace, data)
	}()

	id, err := strconv.Atoi(c.Param("id"))
	if id <= 0 || err != nil {
		retCode = ctrl.ErrRequest
		return
	}

	exist, data, err := model.GetUserById(id)
	if err != nil {
		logger.Log.Error(err.Error())
		retCode = ctrl.ErrInternalServer
		return
	}
	if !exist {
		data = nil
		retCode = ctrl.ErrRequest
		stacktrace = StackTraceIdNotExist
		return
	}
	return
}

// GetUsers doc
// @Summary 获取用户列表
// @Description 获取所有用户
// @Tags
// @Accept       json
// @Produce      json
// @Success 200 {string} string "{"Code":0,"Message":"OK","Data":{}}"
// @Failure 200 {string} string "{"Code":10001,"Message":"请求参数不合法"}"
// @Router /users [get]
func GetUsers(c *gin.Context) {
	retCode := ctrl.OK
	stacktrace := ""
	var data interface{}
	defer func() {
		ctrl.MakeApiRsp(c, retCode, stacktrace, data)
	}()

	data, err := model.GetUsers()
	if err != nil {
		logger.Log.Error(err.Error())
		retCode = ctrl.ErrInternalServer
		return
	}
	return
}
