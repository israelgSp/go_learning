package main

import (
	"testing"
	"reflect"
	"path"
	"runtime"
	"fmt"
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