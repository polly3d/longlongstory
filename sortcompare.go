package main

import (
	"log"
	"math/rand"
	"time"
)

//SortCompare compare all the sort method
//Go benchmark will be better to do this
func SortCompare(st string, total int) {
	if st == "" {
		st = "insertsort"
	}
	arr := make([]int, total)
	for i := 0; i < total; i++ {
		arr[i] = rand.Int()
	}
	defer elapsedTime(st, time.Now())
	switch st {
	case "insertsort":
	case "selectedsort":
	}
}

func elapsedTime(msg string, t time.Time) {
	log.Printf("%s elapsed: %v\n", msg, time.Since(t))
}
