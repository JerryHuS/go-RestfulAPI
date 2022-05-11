/**
 * @Author: alessonhu
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/5/5 12:29
 */

package main

import (
	"apidemo/config"
	"apidemo/logger"
	"apidemo/model"
	"apidemo/router"
)

// @title Swagger API DEMO
// @version v1.0
// @host 127.0.0.1:8888
// @BasePath /api/v1

func main() {
	// 初始化config
	config.InitConfig()

	// 初始化log
	logger.InitLogger()
	defer logger.Sync()

	// 初始化db
	model.InitDB(config.DBConfig.ConnectInfo)

	// 初始化api
	if err := router.NewRouter().Run(config.ServerConfig.HTTPPort); err != nil {
		panic(err)
	}
}
