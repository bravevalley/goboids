package main

func TestBoids(t *testing.T) {

	main()

	if len(boids) != noBoids {
		t.Errorf("The boids in the slice are not up to the number boids wanted")
	}
}
