package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	var largerNumber, temp time.Duration

	var s []time.Duration
	for a := 0; a < 10000; a++ {
		s = append(s, getRequest(int32(a)))

	}
	s = append(s, s[0])
	for _, element := range s {
		if element > temp {
			temp = element
			largerNumber = temp
		}
	}
	fmt.Println("Request cost max :", largerNumber)
}

func getRequest(i int32) time.Duration {

	t := time.Now()
	resp, err := http.Get("https://google.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	elapsed := time.Since(t)
	fmt.Println("tooked", i+1, elapsed)
	return elapsed

}
