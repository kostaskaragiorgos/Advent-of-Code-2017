package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func funcsumdiff() int32 {
	sum := 0

	return int32(sum)
}

func main() {

	b, err := ioutil.ReadFile("day02.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(b), "\n")
	for index, _ := range lines {
		fmt.Println(lines[index])
	}
	fmt.Println(lines[0])

}
