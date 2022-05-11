/**
 * @Author: alessonhu
 * @Description:
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2022/5/5 17:31
 */

package middleware

import "github.com/gin-gonic/gin"

// Auth 鉴权 JWT
func Auth(Option interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ...
		c.Next()
	}
}
