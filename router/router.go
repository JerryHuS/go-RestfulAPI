/**
 * @Author: alessonhu
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2022/5/5 12:52
 */

package router

import (
	"apidemo/config"
	_ "apidemo/docs"
	"apidemo/router/api/users"
	"apidemo/router/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	var r *gin.Engine

	if config.ServerConfig.Mode == config.ModeDebug {
		gin.SetMode(gin.DebugMode)
		r = gin.New()
		r.Use(gin.Logger())
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup := r.Group("/api/v1")
	apiGroup.Use(middleware.Limit(func() {}))
	apiGroup.Use(middleware.Auth(func() {}))
	{
		apiGroup.POST("/users", users.CreateUser)
		apiGroup.GET("/users", users.GetUsers)
		apiGroup.GET("/users/:id", users.GetUserById)
		apiGroup.PUT("/users/:id", users.UpdateUserById)
		apiGroup.DELETE("/users/:id", users.DeleteUserById)
		apiGroup.GET("/nearbyFriends/:name", users.GetNearbyFriendsByName)
	}

	return r
}
