package main

import "math/rand"



type Boid struct {
	Postion vector2D
	Velocity vector2D
	Id int
}

func boidConstructor(id int) *Boid {
	b := Boid{
		Postion: vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		Velocity: vector2D{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
		Id: id,	
	}

	boids[id] = &b
}