package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	Position vector2D
	Velocity vector2D
	Id       int
}

func boidConstructor(id int) {
	b := Boid{
		Position: vector2D{rand.Float64() * (screenWidth * 2), rand.Float64() * (screenHeight * 2)},
		Velocity: vector2D{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
		Id:       id,
	}

	boids[id] = &b

	go b.movement()
}

func (b *Boid) movement() {
	for {
		b.move()
		time.Sleep(time.Millisecond * 5)
	}

}

func (b *Boid) move() {
	b.Position = b.Position.Add(b.Velocity)
	nextPixel := b.Position.Add(b.Velocity)

	if nextPixel.x >= (screenWidth*2) || nextPixel.x < 0 {
		b.Velocity = vector2D{-b.Velocity.x, b.Velocity.y}
	}
	if nextPixel.y >= (screenHeight*2) || nextPixel.y < 0 {
		b.Velocity = vector2D{b.Velocity.x, -b.Velocity.y}
	}

}
