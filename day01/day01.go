package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func captiasum(st string) int32 {
	sum := 0

	for index, _ := range st {
		if index == len(st)-1 {
			if rune(st[index]) == rune(st[0]) {
				number, err := strconv.Atoi(string(st[index]))
				if err != nil {
					log.Fatal(err)
				}
				sum += number

			}
		} else if rune(st[index]) == rune(st[index+1]) {
			number, err := strconv.Atoi(string(st[index]))
			if err != nil {
				log.Fatal(err)
			}
			sum += number
		}
	}
	return int32(sum)
}

func main() {

	b, err := ioutil.ReadFile("day01.txt")
	if err != nil {
		fmt.Println(err)
	}

	st := string(b)
	fmt.Println(captiasum(st))

}
