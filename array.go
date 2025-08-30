package main

import (
	"fmt"
	"strings"
)

func array() {
	// []type{isi}
	// makanan := []string{"Nasi Goreng", "Mie Ayam", "Sate", "Rendang"}
	// fmt.Println(makanan)

	// for i, val := range makanan {
	// 	fmt.Printf("Makanan index ke-%d adalah %s\n", i, val)
	// }

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	_ = numbers
	// loop pada array dengan range
	// for _, val := range numbers {
	// 	fmt.Printf("Hasil kali 3 dari %d adalah %d\n", val, val*3)
	// }

	// mapping
	array := []string{"Nasi Goreng", "Nasi Ayam"}
	array = append(array, "Sate")
	fmt.Println(strings.Join(array, ", "))
	maps := map[int]string{
		1: "Nasi Goreng",
		2: "Mie Ayam",
		3: "Sate",
		4: "Rendang",
	}

	for i, val := range maps {
		fmt.Printf("Makanan  %d adalah %s\n", i, val)
	}

	users := []User{
		{Name: "John", Age: 30, Class: "A"},
		{Name: "Jane", Age: 25, Class: "B"},
		{Name: "Doe", Age: 28, Class: "C"},
	}

	fmt.Println(append(users, User{Name: "Kyy", Age: 16, Class: "A"}))
	// loop pada array dengan index
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Printf("Hasil kali 3 dari %d adalah %d\n", numbers[i], numbers[i]*3)
	// }
}
