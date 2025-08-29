package main

import ( "fmt" )

func variable(){
	var name string = "Fiky Alrasya"
	var ( 
	umur = 16
	Kelas = 12
	Absen = 12
	)
	var (
		y, z = 1, 2
	)
	kanjut := "Woii"
	fmt.Printf("%s, nama saya %s, umur saya %d, saya kelas %d, dan nomor absen saya %d", kanjut, name, umur, Kelas, Absen )
	fmt.Println(y + z)
}
