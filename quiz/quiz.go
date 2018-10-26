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
	"math/rand"
)

type Quiz struct {
	Question, Answer string

}


func check( e error) {
	if e != nil {
		panic(e)
	}

}

//If shuffle flag is set, this is the method
//that shuffles the quiz.
func shuffleQuiz(quiz []Quiz) []Quiz{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(quiz); i>0; i-- {
		j := r.Intn(i)
		quiz[i-1], quiz[j] = quiz[j], quiz[i-1]
	}
	return quiz
}

//This method reads in the .csv file
//and parses into an array of Quiz(es)
func parseCSVFile(filename string) (error, []Quiz) {
	csvFile, err := os.Open(filename)
	check(err)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var quiz []Quiz
	var e error
	for {
		line, e := reader.Read()
		if e == io.EOF {
			break
		}
		quiz = append(quiz, Quiz {
			Question: line[0],
			Answer: line[1],
		})
	}
	return e, quiz
}

//This method does the actual asking of the questions
//and reads the input from the user
func askQuestions(totalCorrectAns * int, quiz []Quiz) {
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

//This the main method the runs the application.
//This method has the timer functionality 
func runQuiz(timerFlag *int, quiz []Quiz) {
	var totalCorrect int
	totalQuestion := len(quiz)
	testFinished := make(chan bool)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please press enter to start quiz: ")
	resp,_ := reader.ReadString('\n')
	resp = strings.TrimSuffix(resp, "\n")
	
	if resp == string('\r') {
		timer := time.NewTimer(time.Duration(*timerFlag) * time.Second).C
		
		go func() {
			askQuestions(&totalCorrect, quiz)
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

func main() { 
	_, filename, _, _ := runtime.Caller(0)
	defaultFilePath := path.Join(path.Dir(filename), "problems.csv")
	filenameFlag := flag.String("filename", defaultFilePath, "filename you would like to pass")
	timerFlag := flag.Int("timer", 30, "time for for timer in seconds")
	shuffleFlag := flag.Bool("shuffle", false, "option to shuffle test")
	flag.Parse()
	var quiz []Quiz
	var e error
	e, quiz = parseCSVFile(*filenameFlag)
	if e != nil {
		panic(e)
	}else if *shuffleFlag {
		quiz = shuffleQuiz(quiz)
	}

	runQuiz(timerFlag, quiz)
	
}