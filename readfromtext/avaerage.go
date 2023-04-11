package main

import (
	"fmt"
	"log"

	"github.com/earthaYan/go_demo/readfromtext/datafile"
)

func average() {
	// D:\private\GoDemo\go_demo\readfromtext\avaerage.go
	numbers, err := datafile.GetFloats("D:\\private\\GoDemo\\go_demo\\readfromtext\\datafile\\data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, val := range numbers {
		sum += val
	}
	count := float64(len(numbers))
	fmt.Println(sum / count)
}
