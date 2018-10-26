package main

import (
	"testing"
	"reflect"
	"path"
	"runtime"
)

//Test shuffle quiz method
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

func TestParseCSVFile(t *testing.T) {
	sizeOfExpectedArray := 12
	_, filename, _, _ := runtime.Caller(0)
	file := path.Join(path.Dir(filename), "problems.csv")
	quiz := parseCSVFile(file)
	if len(quiz) != sizeOfExpectedArray {
		t.Errorf("Size mismatch")
	}

}