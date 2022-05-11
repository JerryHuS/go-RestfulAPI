/**
 * @Author: alessonhu
 * @Description:
 * @File:  geohash
 * @Version: 1.0.0
 * @Date: 2022/5/6 14:11
 */

package utils

import (
	"github.com/mmcloughlin/geohash"
)

// GetGeoHash geohash with the standard 12 characters of precision.
func GetGeoHash(lat, lon float64) string {
	return geohash.Encode(lat, lon)
}
