package main

import (
	"fmt"
	"math"
)

const (
	earthRadius = 6371000 // 地球半径（单位：公里）
)

type Coordinate struct {
	Latitude  float64 // 纬度
	Longitude float64 // 经度
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func calculateDistance(coordinate1, coordinate2 Coordinate) float64 {
	lat1 := degreesToRadians(coordinate1.Latitude)
	lon1 := degreesToRadians(coordinate1.Longitude)
	lat2 := degreesToRadians(coordinate2.Latitude)
	lon2 := degreesToRadians(coordinate2.Longitude)

	deltaLat := lat2 - lat1
	deltaLon := lon2 - lon1

	// 应用Haversine公式计算距离
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c

	return distance
}

func main() {
	// 经纬度坐标点1     113.345786,23.139846
	coordinate1 := Coordinate{Latitude: 23.139846, Longitude: 113.345786}

	// 经纬度坐标点2     113.351881,23.146937
	coordinate2 := Coordinate{Latitude: 23.146937, Longitude: 113.351881}

	// 计算两点之间的距离
	distance := calculateDistance(coordinate1, coordinate2)

	// 打印结果
	fmt.Printf("两点之间的距离为: %.2f米\n", distance)
}
