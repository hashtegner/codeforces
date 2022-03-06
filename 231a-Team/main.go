package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readString() string {
	reader := bufio.NewReader(os.Stdin)
	bytes, _, err := reader.ReadLine()

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

func readInt() int {
	n, err := strconv.Atoi(readString())
	if err != nil {
		log.Fatal(err)
	}

	return n
}

func main() {
	n := readInt()

	count := 0
	for i := 0; i < n; i++ {
		anwser := readString()
		if anwser == "1 0 1" || anwser == "1 1 0" || anwser == "0 1 1" || anwser == "1 1 1" {
			count += 1
		}
	}

	fmt.Println(count)
}
