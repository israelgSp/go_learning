package main

import (
	"io"
	"bufio"
	"os"
	"fmt"
	"encoding/csv"
	"log"
	"strings"
	"flag"
	"runtime"
	"path"
)

type Quiz struct {
	Question, Answer string

}


func check( e error) {
	if e != nil {
		panic(e)
	}

}

func parseCSVFile(filename string) []Quiz {
	csvFile, err := os.Open(filename)
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

func runTest(quiz []Quiz) int {
	var totalCorrectAns int
	reader := bufio.NewReader(os.Stdin)
	for _, quesAns := range quiz {
		question := quesAns.Question
		answer := quesAns.Answer
		fmt.Print(question, ": ")
		response,_ := reader.ReadString('\n')
		response = strings.TrimSuffix(response, "\r\n")
		if response == answer {
			totalCorrectAns += 1
		}
	}

	return totalCorrectAns
}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	defaultFilePath := path.Join(path.Dir(filename), "problems.csv")
	filenameFlag := flag.String("filename", defaultFilePath, "filename you would like to pass")
	flag.Parse()
	var quiz []Quiz
	quiz = parseCSVFile(*filenameFlag)
	totalQuestion := len(quiz)
	totalCorrect := runTest(quiz)
	fmt.Printf("Total correct answers %v out of %v total questions\n", totalCorrect, totalQuestion)

}