/**
 * @Author: alessonhu
 * @Description:
 * @File:  create
 * @Version: 1.0.0
 * @Date: 2022/5/5 14:29
 */

package users

import (
	"apidemo/logger"
	"apidemo/model"
	ctrl "apidemo/router/protocol"
	"apidemo/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	StackTraceIdNotExist = "用户ID不存在"
	StackTraceNameExist  = "用户名已存在"
	StackTraceDobErr     = "日期格式错误"
	StackTraceAddrErr    = "地址格式错误"
)

type UserReq struct {
	Id          int           `json:"id"`
	Name        string        `json:"name"`
	Dob         string        `json:"dob"`
	Addr        model.Address `json:"address"`
	Description string        `json:"description"`
	Following   []int         `json:"following"`
	Followers   []int         `json:"followers"`
}

// CreateUser doc
// @Summary 创建用户
// @Description 创建用户
// @Tags
// @Accept       json
// @Produce      json
// @Param        user  body     UserReq  true  "CreateUser"
// @Success 200 {string} string "{"Code":0,"Message":"OK"}"
// @Failure 200 {string} string "{"Code":10001,"Message":"请求参数不合法"}"
// @Router /users [post]
func CreateUser(c *gin.Context) {
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
		userReq.Addr.GeoHash = utils.GetGeoHash(userReq.Addr.Lat, userReq.Addr.Lon)
		addr, err := json.Marshal(userReq.Addr)
		if err != nil {
			logger.Log.Error(err.Error())
			retCode = ctrl.ErrRequest
			stacktrace = StackTraceAddrErr
			return
		}
		user := &model.User{
			Name:        userReq.Name,
			Dob:         dob,
			Address:     string(addr),
			Description: userReq.Description,
		}
		err = model.InsertUser(user)
		if err != nil {
			if err.Error() == model.UserNameViolate {
				retCode = ctrl.ErrRequest
				stacktrace = StackTraceNameExist
				return
			}
			logger.Log.Error(err.Error())
			retCode = ctrl.ErrInternalServer
			return
		}
		return
	}
	retCode = ctrl.ErrRequest
	return
}
