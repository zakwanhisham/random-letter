package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: randl <target_string>")
		os.Exit(1)
	}

	target_string := os.Args[1]

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ,./?!@#$%^&*()-=_+[]{};'':\"\"|"
	result := ""

	for _, letter := range target_string {
		for {
			l := string(letters[rand.Intn(len(letters))])
			fmt.Println(result + l)

			if l == string(letter) {
				result += l
				break
			}
			time.Sleep(100 * time.Millisecond)
		}

		if result == target_string {
			fmt.Println("---------------")
			fmt.Println(result)
			fmt.Println("---------------")
		}
	}
}
