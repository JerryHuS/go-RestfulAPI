/**
 * @Author: alessonhu
 * @Description:
 * @File:  union
 * @Version: 1.0.0
 * @Date: 2022/5/6 14:38
 */

package utils

func Intersect(slice1, slice2 []int) []int {
	m := make(map[int]int)
	res := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			res = append(res, v)
		}
	}
	return res
}
