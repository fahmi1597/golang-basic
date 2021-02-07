package main

import (
	"fmt"
)

type song struct {
	artist string
	title  string
}

func main() {
	s := song{
		artist: "YUI",
		title:  "Goodbye Days",
	}

	fmt.Printf("(Main func) Now playing: %+v \n", s)
	next(&s)
	fmt.Printf("(Main func) Current song: %+v \n", s)
}

func next(s *song) {
	s.artist = "supercell"
	s.title = "Perfect Day"
	fmt.Printf("(Next func) Now playing: %+v \n", s)
}
