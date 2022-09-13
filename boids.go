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
	// Create a Boid with position and velocity
	b := Boid{
		Position: vector2D{rand.Float64() * (screenWidth * 2), rand.Float64() * (screenHeight * 2)},
		Velocity: vector2D{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
		// Velocity: vector2D{rand.Float64() , rand.Float64()},
		Id:       id,
	}

	// Update the boid slice at index of id with the pointer to the created boid
	boids[id] = &b

	go b.movement()
}

func (b *Boid) movement() {
	// Start and infinite loop that keep calling a move func till the program exits
	for {
		// Move func to test an make the boid move
		b.move()
		time.Sleep(time.Millisecond * 5)
	}

}

func (b *Boid) move() {
	// Update the position of the Boid on the Boid map array before it moves
	boidMap[int(b.Position.x)][int(b.Position.y)] = b.Id
	
	// The new postion of the boid is now the former position plus velocity
	b.Position = b.Position.Add(b.Velocity)
	
	// Update the position of the Boid on the Boid map array after it moves
	boidMap[int(b.Position.x)][int(b.Position.y)] = b.Id

	// Calculate the next postion of the BOID for decision making
	nextPixel := b.Position.Add(b.Velocity)


	// Test for impact against the wall and invert the movement of the boid
	if nextPixel.x >= (screenWidth*2) || nextPixel.x < 0 {
		b.Velocity = vector2D{-b.Velocity.x, b.Velocity.y}
	}
	if nextPixel.y >= (screenHeight*2) || nextPixel.y < 0 {
		b.Velocity = vector2D{b.Velocity.x, -b.Velocity.y}
	}

}
