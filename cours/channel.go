package cours

import (
	"fmt"
	"time"
)

func Master() {
	c := make(chan string)

	go worker(c, func() string { return tache(10) })
	go worker(c, func() string { return tache(5) })
	go worker(c, func() string { return tache(1) })
	go worker(c, func() string { return tache(3) })

	for result := range c {
		fmt.Println(result)
	}
}

func tache(delta time.Duration) string {
	time.Sleep(delta * time.Second)
	return fmt.Sprintf("%v", delta)
}

func worker(c chan string, calback func() string) {
	c <- calback()
}
