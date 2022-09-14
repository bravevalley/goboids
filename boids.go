package main

import (
	"math"
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
		Id: id,
	}

	// Update the boid slice at index of id with the pointer to the created boid
	boids[id] = &b

	boidMap[int(b.Position.x)][int(b.Position.y)] = b.Id

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

func (b *Boid) calcAccer() vector2D {
	var accer vector2D
	avrVel := vector2D{0, 0}
	var count int
	top, lower := b.Position.AddV(float64(boidRadius)), b.Position.AddV(float64(-boidRadius))

	for row := math.Max(lower.x, 0); row <= math.Min(top.x, screenWidth); row++ {
		for col := math.Max(lower.y, 0); col <= math.Min(top.y, screenHeight); col++ {
			boidInView := boidMap[int(row)][int(col)]
			if boidInView != -1 && boidInView != b.Id {
				if boids[boidInView].Position.Distance(b.Position) < boidRadius {
					count++
					avrVel = avrVel.Add(boids[boidInView].Velocity)
				}
			}
		}
	}

	if count > 0 {
		avrVel = avrVel.DivideV(float64(count))
		accer = avrVel.Subtract(avrVel).MultiplyV(perChange)
	}

	return accer
}

func (b *Boid) move() {

	b.Velocity = b.Velocity.Add(b.calcAccer())
	// Update the position of the Boid on the Boid map array before it moves
	boidMap[int(b.Position.x)][int(b.Position.y)] = -1

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
