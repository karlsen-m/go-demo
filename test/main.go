package main

import (
	"math"
)

const (
	// 地球半径，单位：千米
	earthRadius = 6371.0
)

// 计算两点之间的距离
func distance(lat1, lon1, lat2, lon2 float64) float64 {
	// 将经纬度转换为弧度
	lat1 = toRadians(lat1)
	lon1 = toRadians(lon1)
	lat2 = toRadians(lat2)
	lon2 = toRadians(lon2)

	// 使用球面余弦定理计算距离
	deltaLon := lon2 - lon1
	centralAngle := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(deltaLon))
	distance := earthRadius * centralAngle

	return distance
}

// 将角度转换为弧度
func toRadians(angle float64) float64 {
	return angle * (math.Pi / 180.0)
}

func main() {
	// 116.404, 39.915 表示北京天安门广场的经纬度
	// 121.473, 31.230 表示上海外滩的经纬度
	//113.329758, 23.112209
	//113.331303,23.156642
	d := distance(23.112209, 113.329758, 23.156642, 113.331303)
	println(d)
}
