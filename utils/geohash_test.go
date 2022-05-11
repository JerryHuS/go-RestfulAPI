/**
 * @Author: alessonhu
 * @Description:
 * @File:  geohash_test.go
 * @Version: 1.0.0
 * @Date: 2022/5/6 14:14
 */

package utils

import (
	"testing"
)

func TestGetGeoHash(t *testing.T) {
	type args struct {
		lat float64
		lon float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "getGeoHash",
			args: args{
				lat: 118.81874,
				lon: 32.074497,
			},
			want: "hdwccz5utvq2",
		},
		{
			name: "getGeoHash",
			args: args{
				lat: 118.81874,
				lon: 32.174497,
			},
			want: "hdwcuphu8j7b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGeoHash(tt.args.lat, tt.args.lon); got != tt.want {
				t.Errorf("GetGeoHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
