package main

type User struct {
	Name  string
	Age   int
	Class string
}

type Anime struct {
    Title   string   `json:"title"`
    Episode int      `json:"episode"`
    Rating  float64  `json:"rating"`
    Genre   []string `json:"genre"`
}
