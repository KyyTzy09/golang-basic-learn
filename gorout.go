package main

import (
	"fmt"
	"sync"
	"time"
	"encoding/json"
)

func defers() {
	fmt.Println(3)
	defer fmt.Println(2)
	fmt.Println(1)
}

func worker(i int) {
	fmt.Println("Worker 1 :", i)
}

func worker2(i int) {
	fmt.Println("Worker 2 :", i)
}

func push(ch chan string, value string) {
	ch <- value
}

func workers(id int, ch chan string, max int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= max; i++ {
		ch <- fmt.Sprintf("Worker ke %d: %d", id, i)
	}
	time.Sleep(200 * time.Millisecond)
}

func loopChan(max int, ch chan string) {
	for i := 0; i <= max; i++ {
		ch <- fmt.Sprintf("Worker ke-%d", i)
	}
	time.Sleep(200 * time.Millisecond)
}

func withGroupWait(id int, ch chan string, max int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= max; i++ {
		ch <- fmt.Sprintf("Worker ke %d: %d", id, i)
	}
}

func createAnime(data []Anime, ch chan []Anime, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- data
}

func goroutine() {
	var wg sync.WaitGroup
	// Basic
	// fmt.Println("Mulai")
	// var count int = 4
	// for i := 1; i <= count; i++ {
	// 	go worker(i)
	// 	go worker2(i)
	// }
	// time.Sleep(2 * time.Second)
	// fmt.Println("Selesai")

	// With channel
	// fmt.Println("Mulai")
	// var ch = make(chan int, 3)
	// go push(ch, "satu")
	// go push(ch, "dua")
	// time.Sleep(2 * time.Second)
	// ch <- 1
	// ch <- 2
	// ch <- 3

	// fmt.Println(<-ch , <-ch , <-ch)

	// with channel loop
	// fmt.Println("Mulai")
	// max := 4
	// ch := make(chan string, max*3)
	// go loopChan(max, ch)
	// time.Sleep(2 * time.Second)
	// for i := 0; i <= max; i++ {
	// 	fmt.Println(<-ch)
	// }
	// close(ch)
	// v, ok := <-ch
	// if ok {
	// 	fmt.Println("Received:", v)
	// } else {
	// 	fmt.Println("Channel closed")
	// }
	// fmt.Println("Selesai")

	// Latihan 2
	// fmt.Println("Mulai")
	// max := 4
	// ch := make(chan string)
	// for i := 0; i <= 3; i++ {
	// 	wg.Add(1)
	// 	go workers(i, ch, max, &wg)
	// }
	// go func() {
	// 	wg.Wait()
	// 	close(ch)
	// }()

	// for msg := range ch {
	// 	fmt.Println(msg)
	// }

	// with waitgroup
	// fmt.Println("Mulai")
	// ch := make(chan string)
	// max := 5
	// for i := 0; i <= 3; i++ {
	// 	wg.Add(1)
	// 	go withGroupWait( i, ch, max, &wg)
	// }

	// go func() {
	// 	wg.Wait()
	// 	close(ch)
	// }()

	// for msg := range ch {
	// 	fmt.Println(msg)
	// }

	// fmt.Println("Selesai")

	// defers()

	// Percobaan CRUD
	fmt.Println("Mulai")
	ch := make(chan []Anime)
	data := []Anime{{Title: "Yosuga no sora", Episode: 12, Rating: 8.0, Genre: []string{"Ecchi", "Romance"}}, {Title: "Date a live", Episode: 12, Rating: 8.5, Genre: []string{"Action", "Romance", "Harem"}}}

	wg.Add(1)
	go createAnime(data, ch, &wg)
	
	go func() {
		wg.Wait()
		close(ch)
	}()

	for anime := range ch {
		d, err := json.Marshal(anime)
		if (err != nil) {
			fmt.Println("Error:", err)
		}
		fmt.Println(string(d))
	}
}
