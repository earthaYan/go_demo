package main

import "fmt"

func main() {
	notes := make([]string, 11)
	notes[0] = "do"
	notes = append(notes, "re")
	fmt.Println(notes, len(notes), cap(notes))
}
