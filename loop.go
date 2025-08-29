package main

import "fmt"

func loop() {
	// for loop
	for count:= 0; count <= 100; count ++ {
		fmt.Printf("Perulangan ke-%d\n", count)
	}

	// while loop

	c1 := 1
	max := 100
	for c1 <= max {
		fmt.Printf("Perulangan ke-%d\n", c1)
		c1++
	}

	// loop with conditional
	for counter := 1; counter <= 100; counter++ {
		if counter%2 == 0 {
			fmt.Printf("bilangan genap perulangan ke %d\n", counter)
		} else {
			fmt.Printf("bilangan ganjil perulangan ke %d\n", counter)
		}
	}
}
