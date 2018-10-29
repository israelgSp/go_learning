package main

import (
	"testing"
	"reflect"
	"path"
	"runtime"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//Test shuffle quiz method. Test to see if file is shuffled correctly
func TestShuffle(t *testing.T) {
	quiz := []Quiz{
		{"5+5","10"},
		{"1+1","2"},
		{"8+3","11"},
		{"1+2","3"},
		{"8+6","14"},
	}
	shuffledQuiz := shuffleQuiz(quiz)
	return_bool := reflect.DeepEqual(quiz, shuffledQuiz)
	
	if return_bool == false {
		t.Errorf("Quiz has not been shuffled")
	}

	quiz1 := []Quiz{
		{"Who is the president of the united states?","Trump"},
		{"What is the captial of USA?","D.C."},
		{"What is Israel's favorite color?","Green"},
		{"What is the captial of Colorado?","Denver"},
		{"What is the professional basketball organization in USA?","NBA"},
	}

	shuffledQuiz1 := shuffleQuiz(quiz1)
	return_bool1 := reflect.DeepEqual(quiz1, shuffledQuiz1)

	if return_bool1 == false {
		t.Errorf("Quiz has not been shuffled")
	}

}

//Test if method parses csv file correctly
func TestParseCSVFileParsesFileCorrectly(t *testing.T) {
	sizeOfExpectedArray := 12
	_, filename, _, _ := runtime.Caller(0)
	file := path.Join(path.Dir(filename), "problems.csv")

	_, quiz := parseCSVFile(file)
	if len(quiz) != sizeOfExpectedArray {
		t.Errorf("Size mismatch")
	}

	sizeOfExpectedArray1 := 5
	_, filename1, _, _ := runtime.Caller(0)
	file1 := path.Join(path.Dir(filename1), "random.csv")

	_, quiz1 := parseCSVFile(file1)
	if len(quiz1) != sizeOfExpectedArray1 {
		t.Errorf("Size mismatch")
	}
}

//Test if method handles not having an existing file correctly
func TestParseCSVFileCorrectErrorIfFileDoesNOTExits(t *testing.T) {
	notRealFile := "not real file"
	e,_ := parseCSVFile(notRealFile)

	if e.Error() != fmt.Sprintf("open %s: The system cannot find the file specified.", notRealFile) {
		t.Errorf("Expected 'The system cannot find the file specified' error")
	}

	notRealFile1 := "randoms.csv"
	e1,_ := parseCSVFile(notRealFile1)

	if e1.Error() != fmt.Sprintf("open %s: The system cannot find the file specified.", notRealFile1) {
		t.Errorf("Expected 'The system cannot find the file specified' error")
	}
}

//Testing askQuestions method to see if method calculates correct answers correctly.
func TestAskQuestionsCalculatesAnswersCorrectly(t *testing.T) {
	var totalCorrectCalculated int
	totalCorrect := 5
	//user input setup
	in, err := ioutil.TempFile("", "")

	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()
	
	//user input to test. All correct
	_, err = io.WriteString(in, "10\n2\n11\n3\n14\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)

	quiz := []Quiz{
		{"5+5","10"},
		{"1+1","2"},
		{"8+3","11"},
		{"1+2","3"},
		{"8+6","14"},
	}
	
	askQuestions(in, &totalCorrectCalculated, quiz)

	if totalCorrectCalculated != totalCorrect {
		t.Errorf("Correct answers were not calcuated properly")
	}

	//user input to test with 4 correct answers

	in, err = ioutil.TempFile("", "")

	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = io.WriteString(in, "10\n2\n11\n3\n1\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)

	totalCorrect = 4
	totalCorrectCalculated = 0


	askQuestions(in, &totalCorrectCalculated, quiz)

	if totalCorrectCalculated != totalCorrect {
		t.Errorf("Correct answers were not calcuated properly")
	}

	//user input to test with an invalid answer 4 correct answers

	in, err = ioutil.TempFile("", "")

	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = io.WriteString(in, "10\n2\n11\nrandom\n14\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)

	totalCorrect = 4
	totalCorrectCalculated = 0


	askQuestions(in, &totalCorrectCalculated, quiz)

	if totalCorrectCalculated != totalCorrect {
		t.Errorf("Correct answers were not calcuated properly")
	}
}