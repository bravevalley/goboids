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
		Position: vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
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
	var avrPos vector2D
	var sep vector2D

	// Lock the data to be read and written to - boidRadius, BoidMap,
	// the boids field; the captured boid might have a thread writing
	// to it by modifying its fields eg position
	rWlock.RLock()
	top, lower := b.Position.AddV(float64(boidRadius)), b.Position.AddV(float64(-boidRadius))

	for row := math.Max(lower.x, 0); row <= math.Min(top.x, screenWidth); row++ {
		for col := math.Max(lower.y, 0); col <= math.Min(top.y, screenHeight); col++ {
			boidInView := boidMap[int(row)][int(col)]
			if boidInView != -1 && boidInView != b.Id {
				if dist := boids[boidInView].Position.Distance(b.Position); dist < boidRadius {
					count++
					avrVel = avrVel.Add(boids[boidInView].Velocity)
					avrPos = avrPos.Add(boids[boidInView].Position)
					sep = sep.Add(b.Position.Subtract(boids[boidInView].Position).DivideV(dist))
				}
			}
		}
	}
	rWlock.RUnlock()

	accer = vector2D{b.borderBounce(b.Position.x, screenWidth), b.borderBounce(b.Position.y, screenHeight)}
	if count > 0 {
		avrVel = avrVel.DivideV(float64(count)).Subtract(b.Velocity).MultiplyV(perChange)
		avrPos = avrPos.DivideV(float64(count)).Subtract(b.Position).MultiplyV(perChange)
		sepeRatio := sep.MultiplyV(perChange)
		accer = accer.Add(avrVel).Add(avrPos).Add(sepeRatio)
	}
	return accer
}

func (b *Boid) move() {
	// New accerlation value is placed outside the Mutex to prevent deadlock
	acceleration := b.calcAccer()

	// Lock the data to be read and written to - boidRadius, BoidMap,
	rWlock.Lock()
	b.Velocity = b.Velocity.Add(acceleration).limitV(1, -1)
	// Update the position of the Boid on the Boid map array before it moves
	boidMap[int(b.Position.x)][int(b.Position.y)] = -1

	// The new postion of the boid is now the former position plus velocity
	b.Position = b.Position.Add(b.Velocity)

	// Update the position of the Boid on the Boid map array after it moves
	boidMap[int(b.Position.x)][int(b.Position.y)] = b.Id

	// unlock the data to be read and written so other thread can use it as well
	rWlock.Unlock()
}

func (b *Boid) borderBounce(position, maxBorderPos float64) float64 {

	// Test if position x or y is less that the view radius of the boid,
	// this means the boid is close to origin
	if position < boidRadius {
		// change the velocity to the reciprocal of the position ie 1 / position
		// e.g 1 / 9 = 0.9 the new velocity of the boid
		// note you can only be on a negative velocity to traverse back to the
		// origin because the p-rogram init with positve velocity meaning
		// you are moving away from the origin.
		return 1 / position
	} else if position > maxBorderPos-boidRadius {
		// 1 / 632 - 640
		// 1 / -8 == -0.8
		// This way you are mnoving away from the limit
		return 1 / (position - maxBorderPos)
	}

	// Unreachable
	return 0
}
