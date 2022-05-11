/**
 * @Author: alessonhu
 * @Description:
 * @File:  friends
 * @Version: 1.0.0
 * @Date: 2022/5/6 12:58
 */

package users

import (
	"apidemo/logger"
	"apidemo/model"
	ctrl "apidemo/router/protocol"
	"github.com/gin-gonic/gin"
)

// GetNearbyFriendsByName doc
// @Summary 获取附近朋友
// @Description 获取附近朋友
// @Tags
// @Accept       json
// @Produce      json
// @Param        name  path  string  true  "Username"
// @Success 200 {string} string "{"Code":0,"Message":"OK","Data":{}}"
// @Failure 200 {string} string "{"Code":10001,"Message":"请求参数不合法"}"
// @Router /nearbyFriends/{name} [get]
func GetNearbyFriendsByName(c *gin.Context) {
	retCode := ctrl.OK
	stacktrace := ""
	var data interface{}
	defer func() {
		ctrl.MakeApiRsp(c, retCode, stacktrace, data)
	}()

	name := c.Param("name")
	if name == "" {
		retCode = ctrl.ErrRequest
		return
	}
	data, err := model.GetNearByFriendsByName(name)
	if err != nil {
		logger.Log.Error(err.Error())
		retCode = ctrl.ErrInternalServer
		return
	}
	return
}
