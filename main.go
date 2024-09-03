package main

import (
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

	fmt.Println("Press Enter to start Quiz")

	startQuiz(records)

	// fmt.Printf("Questions answered correctly: %v\n", correctAnswers)
	// fmt.Printf("Questions answered incorrectly: %v\n", incorrectAnswers)

}

func startQuiz(records [][]string) {

	var correctAnswers []string
	var incorrectAnswers []string

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
		fmt.Printf("Questions answered correctly: %v\n", correctAnswers)
		fmt.Printf("Questions answered incorrectly: %v\n", incorrectAnswers)
		os.Exit(2)
	case <-time.After(time.Second * 60):
		fmt.Printf("Questions answered correctly: %v\n", correctAnswers)
		fmt.Printf("Questions answered incorrectly: %v\n", incorrectAnswers)
		os.Exit(1)

	}

	// return correctAnswers, incorrectAnswers

}
