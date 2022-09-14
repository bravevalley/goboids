package main

import "math"

type vector2D struct {
	x float64
	y float64
}

// Add two vectors
func (v1 vector2D) Add(v2 vector2D) vector2D {
	return vector2D{v1.x + v2.x, v1.y + v2.y}
}

// Subtract two vectors
func (v1 vector2D) Subtract(v2 vector2D) vector2D {
	return vector2D{v1.x - v2.x, v1.y - v2.y}
}

func (v1 vector2D) AddV(value float64) vector2D {
	return vector2D{v1.x + value, v1.y + value}
}

func (v1 vector2D) DivideV(value float64) vector2D {
	return vector2D{v1.x / value, v1.y / value}
}

func (v1 vector2D) MultiplyV(value float64) vector2D {
	return vector2D{v1.x * value, v1.y * value}
}

func (v1 vector2D) Distance(v2 vector2D) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
