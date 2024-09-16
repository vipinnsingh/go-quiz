package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {

	file, err := os.Open("problems.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Press Enter to start Quiz")

	bufReader := bufio.NewReader(os.Stdin)

	_, _ = bufReader.ReadString('\n')

	fmt.Println("Lets start the Quiz!!")
	startQuiz(records)

	// fmt.Printf("Questions answered correctly: %v\n", correctAnswers)
	// fmt.Printf("Questions answered incorrectly: %v\n", incorrectAnswers)

}

func startQuiz(records [][]string) {

	var correctAnswers []string
	var incorrectAnswers []string

	fmt.Print("Please enter the duration (in seconds) for test :")
	var timer int
	fmt.Scanln(&timer)

	c := make(chan bool)

	go func() {
		for i := 0; i < len(records); i++ {
			var answer string
			fmt.Printf("What is %v ?", records[i][0])
			fmt.Scanln(&answer)
			if answer == records[i][1] {
				correctAnswers = append(correctAnswers, records[i][0])
			} else {
				incorrectAnswers = append(incorrectAnswers, records[i][0])
			}
		}
		c <- true
	}()

	select {

	case <-c:
		fmt.Println()
		fmt.Printf("Questions answered correctly: %v\n", correctAnswers)
		fmt.Printf("Questions answered incorrectly: %v\n", incorrectAnswers)
		break
	case <-time.After(time.Second * time.Duration(timer)):
		fmt.Println()
		fmt.Printf("Questions answered correctly: %v\n", correctAnswers)
		fmt.Printf("Questions answered incorrectly: %v\n", incorrectAnswers)
		break

	}

	// return correctAnswers, incorrectAnswers

}
