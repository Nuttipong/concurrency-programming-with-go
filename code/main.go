package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		/* Simple func */
		// if b, ok := queryCache(id); ok {
		// 	fmt.Println("from cache")
		// 	fmt.Println(b)
		// 	continue
		// }
		// if b, ok := queryDatabase(id); ok {
		// 	fmt.Println("from database")
		// 	cache[id] = b
		// 	fmt.Println(b)
		// 	continue
		// }
		// fmt.Printf("Book not found id: '%v'", id)

		/* Goroutines func */
		go func(id int) {
			if b, ok := queryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
		}(id)
		go func(id int) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				cache[id] = b
				fmt.Println(b)
			}
		}(id)
		/* abilities to setup the cache */
		time.Sleep(50 * time.Millisecond)
	}
	/* abilities to wait until goroutines finished the task. just ephemeral way*/
	time.Sleep(3 * time.Second)
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(300 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			return b, true
		}
	}

	return Book{}, false
}
