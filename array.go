package main

import "fmt"

func array() {
	// []type{isi}
	makanan := []string{"Nasi Goreng", "Mie Ayam", "Sate", "Rendang"}
	fmt.Println(makanan)

	for i, val := range makanan {
		fmt.Printf("Makanan index ke-%d adalah %s\n", i, val)
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// loop pada array dengan range
	for _, val := range numbers {
		fmt.Printf("Hasil kali 3 dari %d adalah %d\n", val, val*3)
	}

	// loop pada array dengan index
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Hasil kali 3 dari %d adalah %d\n", numbers[i], numbers[i]*3)
	}
}
