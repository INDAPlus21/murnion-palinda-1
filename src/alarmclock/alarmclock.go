package main

import (
	"fmt"
	"time"
)

func Remind(text string, seconds int64) {
	for {
		t := time.Now().Local().Format("The time is 15.04.05:")
		fmt.Println(t, text)
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}

func main() {
	go Remind("Time to eat", 10)
	go Remind("Time to work", 30)
	go Remind("Time to sleep", 60)
	select {}
}
