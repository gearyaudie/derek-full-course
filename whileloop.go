// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// var pl = fmt.Println

// func main() {

// 	aNums:= []int{1,2,3}

// 	for _, num := range aNums {
// 		pl(num)
// 	}

// 	// seedSecs := time.Now().Unix()
// 	// rand.Seed(seedSecs)
// 	// randNum := rand.Intn(50) + 1
// 	// for true {
// 	// 	fmt.Println("Guess a number between 0 and 50: ")
// 	// 	pl("Random number is: ", randNum)
// 	// 	reader := bufio.NewReader(os.Stdin)
// 	// 	guess, err := reader.ReadString('\n')
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// 	guess = strings.TrimSpace(guess)
// 	// 	iGuess, err := strconv.Atoi(guess)
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// 	if iGuess > randNum {
// 	// 		pl("Lower")
// 	// 	} else if iGuess < randNum {
// 	// 		pl("Pick a higher value")
// 	// 	} else {
// 	// 		pl("You guessed it")
// 	// 		break
// 	// 	}
// 	// }

// }
