/**
 * @Author: alessonhu
 * @Description:
 * @File:  db
 * @Version: 1.0.0
 * @Date: 2022/5/5 11:21
 */

package model

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"time"
	"xorm.io/core"
)

var (
	Db *xorm.Engine
)

func InitDB(psqlInfo string) error {
	var err error
	Db, err = xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = Db.Ping(); err != nil {
		panic(err)
	}
	Db.SetMaxIdleConns(10)
	Db.SetMaxOpenConns(20)
	Db.SetConnMaxLifetime(time.Minute * 5)
	Db.SetLogLevel(core.LOG_INFO)
	go keepDbAlived(Db)
	return err
}

func keepDbAlived(engine *xorm.Engine) {
	t := time.Tick(time.Minute * 3)
	for {
		<-t
		engine.Ping()
	}
}
