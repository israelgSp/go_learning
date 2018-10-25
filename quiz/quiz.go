package main

import (
	"io"
	"bufio"
	"os"
	"fmt"
	"encoding/csv"
	"strings"
	"flag"
	"runtime"
	"path"
	"time"
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
		} else  {
			check(error)
		}
		quiz = append(quiz, Quiz {
			Question: line[0],
			Answer: line[1],
		})
	}
	return quiz
}

func runTest(totalCorrectAns * int, quiz []Quiz) {
	reader := bufio.NewReader(os.Stdin)
	for _, quesAns := range quiz {
		question := quesAns.Question
		answer := quesAns.Answer
		fmt.Print(question, ": ")
		response,_ := reader.ReadString('\n')
		response = strings.TrimSuffix(response, "\r\n")
		response = strings.TrimSpace(response)
		if strings.EqualFold(response, answer) {
			*totalCorrectAns += 1
		}
	}
}

func main() { 
	_, filename, _, _ := runtime.Caller(0)
	defaultFilePath := path.Join(path.Dir(filename), "problems.csv")
	filenameFlag := flag.String("filename", defaultFilePath, "filename you would like to pass")
	flag.Parse()
	var quiz []Quiz
	var totalCorrect int
	testFinished := make(chan bool)
	quiz = parseCSVFile(*filenameFlag)
	totalQuestion := len(quiz)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please press enter to start quiz: ")
	resp,_ := reader.ReadString('\n')
	resp = strings.TrimSuffix(resp, "\n")
	
	if resp == string('\r') {
		timer := time.NewTimer(30 * time.Second).C
		
		go func() {
			runTest(&totalCorrect, quiz)
			testFinished <- true
		}()
		
		for {
			select {
			case <- timer:
				fmt.Println("time expired")
				fmt.Printf("Total correct answers %v out of %v total questions\n", totalCorrect, totalQuestion)
				return
			case <- testFinished:
				fmt.Printf("Total correct answers %v out of %v total questions\n", totalCorrect, totalQuestion)
				return
			}
	
		}
	
	}
	
}