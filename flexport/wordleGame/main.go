package main

import (
	"fmt"
	"strings"
)

func getResultByGuess(target string, guess string) string {
	result := ""
	for i := 0; i < len(guess); i++ {
		if target[i] == guess[i] {
			result += "G"
		} else if strings.Contains(target, string(guess[i])) {
			result += "Y"
		} else {
			result += "B"
		}
	}
	return result
}

func getSuitableGuessByGuessList(guessList []string, target string, output string) []string {
	suitableGuseeList := []string{}
	for _, guess := range guessList {
		if getResultByGuess(target, guess) == output {
			suitableGuseeList = append(suitableGuseeList, guess)
		}
	}
	return suitableGuseeList
}

func main() {
	target := ""
	guess := ""

	fmt.Print("target: ")
	fmt.Scanln(&target)

	fmt.Print("guess: ")
	fmt.Scanln(&guess)

	result := getResultByGuess(target, guess)
	fmt.Println("result: ", result)

}
