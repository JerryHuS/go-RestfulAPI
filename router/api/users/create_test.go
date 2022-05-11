/**
 * @Author: alessonhu
 * @Description:
 * @File:  create_test.go
 * @Version: 1.0.0
 * @Date: 2022/5/6 13:02
 */

package users

import (
	"apidemo/model"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	r := SetupRouter()
	r.POST("/users", CreateUser)
	//创建post请求体
	reqBody := UserReq{
		Id:          0,
		Name:        "alessonhu",
		Dob:         "2020-01-01",
		Addr:        model.Address{},
		Description: "hello",
		Following:   []int{},
		Followers:   []int{},
	}
	//序列化请求体
	jsonValue, _ := json.Marshal(reqBody)
	//发起post请求
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}
