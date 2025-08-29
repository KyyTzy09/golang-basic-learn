package main
import "fmt"

func conditional(){
	// if else
	hari := "Senin"
	if hari == "Senin" {
		fmt.Printf("Ini hari %s, warung belum buka", hari)
	} else if hari == "Selasa" {
		fmt.Printf("Ini hari %s, warung buka", hari)
	} else {
		fmt.Printf("Ini hari %s, warung sedang libur", hari)
	}

	// switch case
	switch hari {
		case "senin" : {
			fmt.Printf("Ini hari %s, warung buka", hari)
		}
		case "selasa" : {
			fmt.Printf("Ini hari %s, warung buka", hari)
		}
		default: {
			fmt.Printf("Ini hari %s, warung sedang libur", hari)
		}
	}
}