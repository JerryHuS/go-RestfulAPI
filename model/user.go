/**
 * @Author: alessonhu
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/5/5 15:26
 */

package model

import (
	"apidemo/utils"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type User struct {
	Id          int       `xorm:"not null pk autoincr unique INTEGER"`
	Name        string    `xorm:"VARCHAR(255)"`
	Dob         time.Time `xorm:"DATE"`
	Address     string    `xorm:"JSONB"`
	Description string    `xorm:"text"`
	Following   IntArray  `xorm:"text"`
	Followers   IntArray  `xorm:"text"`
	Itime       time.Time `xorm:"created"`
	Utime       time.Time `xorm:"updated"`
}

type Address struct {
	City    string  `json:"city"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	GeoHash string  `json:"geohash"`
}

const (
	UserNameViolate = `pq: duplicate key value violates unique constraint "user_name_key"`
)

func InsertUser(user *User) (err error) {
	if user.Address == "" {
		user.Address = "{}"
	}
	_, err = Db.Insert(user)
	return
}

func GetUserById(id int) (exist bool, user User, err error) {
	exist, err = Db.Id(id).Get(&user)
	return
}

func GetUsers() (user []User, err error) {
	err = Db.Find(&user)
	return
}

func UpdateUserById(user *User) (affected int64, err error) {
	if user.Address == "" {
		user.Address = "{}"
	}
	affected, err = Db.Id(user.Id).Update(user)
	return
}

func DeleteUserById(id int) (affected int64, err error) {
	user := new(User)
	affected, err = Db.Id(id).Delete(user)
	return
}

func GetNearByFriendsByName(name string) (res []User, err error) {
	var user User
	exist, err := Db.Where("name = ?", name).Get(&user)
	if err != nil {
		return
	}
	if !exist {
		return
	}

	// 互关就是朋友
	friendsIds := utils.Intersect(user.Following, user.Followers)
	friendsIdsInStr := strings.Replace(strings.Trim(fmt.Sprint(friendsIds), "[]"), " ", ",", -1)
	if len(friendsIds) == 0 {
		return
	}

	var addr Address
	err = json.Unmarshal([]byte(user.Address), &addr)
	// 没有geohash就不计算，返回空
	if err != nil || len(addr.GeoHash) != 12 {
		return
	}
	geohashPre := addr.GeoHash[0:6]

	whereSql := fmt.Sprintf("id in (%s) and address::jsonb->>'geohash' like '%s%%'", friendsIdsInStr, geohashPre)
	fmt.Println(whereSql)
	err = Db.Where(whereSql).Find(&res)
	if err != nil {
		return
	}
	return
}
