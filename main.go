package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func input(wg *sync.WaitGroup, ch1 chan string) {
	defer wg.Done()
	filepath := "1.txt"
	file, err := os.ReadFile(filepath)
	if err != nil {
		return
	}
	split := strings.Split(string(file), "\r\n")
	for _, v := range split {
		ch1 <- v
	}
	close(ch1)
}
func shift(wg *sync.WaitGroup, ch1 chan string, ch2 chan []string) {
	defer wg.Done()
	for {
		s, ok := <-ch1
		if !ok {
			close(ch2)
			return
		} else {
			split := strings.Split(s, " ")
			n := len(split)
			for i := 1; i <= n; i++ {
				split = append(split, split[0])
				split = split[1:]
				ch2 <- split
			}
		}
	}
}
func alphabetizer(wg *sync.WaitGroup, ch2 chan []string, ch3 chan []string) {
	defer wg.Done()
	for {
		s, ok := <-ch2
		if !ok {
			close(ch3)
			return
		} else {
			for k, v := range s {
				s[k] = strings.ToLower(v)
			}
			ch3 <- s
		}
	}
}
func output(wg *sync.WaitGroup, ch3 chan []string) {
	defer wg.Done()
	for {
		s, ok := <-ch3
		if !ok {
			return
		} else {
			for _, v := range s {
				fmt.Printf("%s ", v)
			}
			fmt.Println()
		}
	}
}
func main() {
	wg := sync.WaitGroup{}
	ch1 := make(chan string)
	ch2 := make(chan []string)
	ch3 := make(chan []string)
	wg.Add(4)
	go input(&wg, ch1)
	go shift(&wg, ch1, ch2)
	go alphabetizer(&wg, ch2, ch3)
	go output(&wg, ch3)
	wg.Wait()
}
