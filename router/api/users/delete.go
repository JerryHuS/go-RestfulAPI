/**
 * @Author: alessonhu
 * @Description:
 * @File:  delete
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

// DeleteUserById doc
// @Summary 删除用户
// @Description 根据id删除指定用户
// @Tags
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Success 200 {string} string "{"Code":0,"Message":"OK"}"
// @Failure 200 {string} string "{"Code":10001,"Message":"请求参数不合法"}"
// @Router /users/{id} [delete]
func DeleteUserById(c *gin.Context) {
	retCode := ctrl.OK
	stacktrace := ""
	defer func() {
		ctrl.MakeApiRsp(c, retCode, stacktrace, nil)
	}()

	id, err := strconv.Atoi(c.Param("id"))
	if id <= 0 || err != nil {
		retCode = ctrl.ErrRequest
		return
	}

	affected, err := model.DeleteUserById(id)
	if affected == 0 {
		retCode = ctrl.ErrRequest
		stacktrace = StackTraceIdNotExist
		return
	}
	if err != nil {
		logger.Log.Error(err.Error())
		retCode = ctrl.ErrInternalServer
		return
	}
	return
}
