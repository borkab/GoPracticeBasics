package main

import "math"

type Square struct {
	Side float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func Info(sha Shape) float64 {
	return sha.Area()
}
