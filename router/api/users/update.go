/**
 * @Author: alessonhu
 * @Description:
 * @File:  update
 * @Version: 1.0.0
 * @Date: 2022/5/5 17:26
 */

package users

import (
	"apidemo/logger"
	"apidemo/model"
	ctrl "apidemo/router/protocol"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

// UpdateUserById doc
// @Summary 更新用户
// @Description 更新用户信息
// @Tags
// @Accept       json
// @Produce      json
// @Param        user  body     UserReq  true  "CreateUser"
// @Success 200 {string} string "{"Code":0,"Message":"OK"}"
// @Failure 200 {string} string "{"Code":10001,"Message":"请求参数不合法"}"
// @Router /users/{id} [put]
func UpdateUserById(c *gin.Context) {
	var userReq UserReq
	retCode := ctrl.OK
	stacktrace := ""
	defer func() {
		ctrl.MakeApiRsp(c, retCode, stacktrace, nil)
	}()

	if err := c.ShouldBind(&userReq); err == nil {
		dob, err := time.Parse("2006-01-02", userReq.Dob)
		if err != nil {
			logger.Log.Error(err.Error())
			retCode = ctrl.ErrRequest
			stacktrace = StackTraceDobErr
			return
		}
		addr, err := json.Marshal(userReq.Addr)
		if err != nil {
			logger.Log.Error(err.Error())
			retCode = ctrl.ErrRequest
			stacktrace = StackTraceAddrErr
			return
		}
		user := &model.User{
			Id:          userReq.Id,
			Name:        userReq.Name,
			Dob:         dob,
			Address:     string(addr),
			Description: userReq.Description,
			Following:   userReq.Following,
			Followers:   userReq.Followers,
		}
		affected, err := model.UpdateUserById(user)
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
	retCode = ctrl.ErrRequest
	return
}
