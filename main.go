package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

func check (min, max int) {
	maps := make(map[[16]byte]bool)
	for i := min; i < max; i++ {
		hash := md5.Sum([]byte(strconv.Itoa(i)))
		_, ok := maps[hash]
		if ok {			
			fmt.Printf("collision!")
			break
		} else {
			maps[hash] = true
			fmt.Printf("\r%d", i)
		}
	}
	fmt.Println()
}
func main() {
	fmt.Println("start")
	start := time.Now()
	check(100000000, 110000000)
	fmt.Println("end")
	end := time.Now()
	fmt.Println(end.Sub(start))
} 