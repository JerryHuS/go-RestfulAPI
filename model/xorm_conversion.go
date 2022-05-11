/**
 * @Author: alessonhu
 * @Description:
 * @File:  xorm_conversion
 * @Version: 1.0.0
 * @Date: 2022/5/5 15:36
 */

package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type IntArray []int

func (s *IntArray) FromDB(bts []byte) error {
	if len(bts) == 0 {
		return nil
	}

	str := string(bts)
	if strings.HasPrefix(str, "{") {
		str = "[" + str[1:len(str)]
	}

	if strings.HasSuffix(str, "}") {
		str = str[0:len(str)-1] + "]"
	}

	var ia = &[]int{}

	err := json.Unmarshal([]byte(str), ia)
	if err != nil {
		return err
	}

	*s = IntArray(*ia)
	return nil
}

func (s *IntArray) ToDB() ([]byte, error) {
	return serializeBigIntArray(*s, "{", "}"), nil
}

func (arr IntArray) MarshalJSON() ([]byte, error) {
	return serializeBigIntArrayAsString(arr, "[", "]"), nil
}

func (arr *IntArray) UnmarshalJSON(b []byte) error {
	var strarr []string
	var intarr []int

	err := json.Unmarshal(b, &strarr)
	if err != nil {
		return err
	}

	for _, s := range strarr {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}

		intarr = append(intarr, i)
	}

	*arr = intarr
	return nil
}

func serializeBigIntArray(s []int, prefix string, suffix string) []byte {
	var buffer bytes.Buffer

	buffer.WriteString(prefix)

	for idx, val := range s {
		if idx > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(strconv.Itoa(val))
	}

	buffer.WriteString(suffix)

	return buffer.Bytes()
}

func serializeBigIntArrayAsString(s []int, prefix string, suffix string) []byte {
	var buffer bytes.Buffer

	buffer.WriteString(prefix)

	for idx, val := range s {
		if idx > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString("\"")
		buffer.WriteString(strconv.Itoa(val))
		buffer.WriteString("\"")
	}
	buffer.WriteString(suffix)
	return buffer.Bytes()
}

type JsonTime time.Time

// 实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

//type JsonTime time.Time

func (j JsonTime) format() string {
	return time.Time(j).Format("2006-01-02 15:04:05")
}

func (j JsonTime) MarshalText() ([]byte, error) {
	return []byte(j.format()), nil
}
