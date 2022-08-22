package main

import (
	"fmt"
	"math"
)

/*
TODO Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными
параметрами x,y и конструктором.
*/

const (
	PI           = 3.1415926535
	EARTH_RADIUS = 6371000
)

type point struct {
	latitude, longitude float64
}

func NewPoint(latitude, longitude float64) *point {
	return &point{latitude, longitude}
}

func (p *point) convertDegreesToRadians() point {
	return point{latitude: p.latitude * PI / 180, longitude: p.longitude * PI / 180}
}

func (p *point) DistanceToPoint(otherP *point) float64 {
	lhs := p.convertDegreesToRadians()
	rhs := otherP.convertDegreesToRadians()

	return math.Acos(
		math.Sin(lhs.latitude)*math.Sin(rhs.latitude)+
			math.Cos(lhs.latitude)*math.Cos(rhs.latitude)*
				math.Cos(math.Abs(lhs.longitude-rhs.longitude)),
	) *
		EARTH_RADIUS
}

func main() {
	p1 := NewPoint(23, 21)
	p2 := NewPoint(43, 21)
	fmt.Println(p1.DistanceToPoint(p2))
}
