/**
 * @Author: alessonhu
 * @Description:
 * @File:  limit
 * @Version: 1.0.0
 * @Date: 2022/5/5 17:31
 */

package middleware

import (
	"github.com/gin-gonic/gin"
)

// Limit 限流熔断...
func Limit(Option interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ...
		c.Next()
	}
}
