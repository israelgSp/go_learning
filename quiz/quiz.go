package main

import (
	"io"
	"bufio"
	"os"
	"fmt"
	"encoding/csv"
	"log"
)

type Quiz struct {
	Question, Answer string

}


func check( e error) {
	if e != nil {
		panic(e)
	}

}

func parseCSVFile() []Quiz {
	csvFile, err := os.Open("c:/Users/israelg/Documents/go_learning/quiz/problems.csv")
	check(err)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var quiz []Quiz
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		quiz = append(quiz, Quiz {
			Question: line[0],
			Answer: line[1],
		})
	}
	return quiz
}

func runTest(quiz []Quiz)  (int, int) {
	reader := bufio.NewReader(os.Stdin)
	for _, quesAns := range quiz {
		question := quesAns.Question
		answer := quesAns.Answer
		fmt.Print(quesAns)
	}


}

func main() {
	var quiz []Quiz
	quiz = parseCSVFile()



}