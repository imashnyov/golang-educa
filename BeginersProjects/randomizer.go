/*
Lucky number game
*/
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 4 //game difficulty
	usage    = `
Welcome to the Lucky Number game!
Pick a number as an argument`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args1, args2 := os.Args[1:2]
	if len(args) != 1 {
		fmt.Printf("%s\n", usage)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		fmt.Println(n, guess)
		if n == guess && turn == 0 {
			fmt.Printf("\t\t🎉 EXTRA WIN 🎉\n \t(victory on the first step) \n")
			return
		}
		if n == guess {
			switch rand.Intn(3) {
			case 0:
				fmt.Println("🎉  YOU WIN!")
			case 1:
				fmt.Println("🎉  YOU'RE AWESOME!")
			case 2:
				fmt.Println("🎉  PERFECT!")
			}
			return
		}
	}

	if guess < 0 {
		fmt.Println("Pick a positive number.")
		return
	}

	fmt.Println("You lost... Try again?")
}
