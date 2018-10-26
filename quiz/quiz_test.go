package main

import (
	"testing"
	"runtime"
	"path"
	"reflect"
)

//Test shuffle quiz method
func TestShuffle(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	file := path.Join(path.Dir(filename), "problems.csv")
	quiz := parseCSVFile(file)
	shuffledQuiz := shuffleQuiz(quiz)
	return_bool := reflect.DeepEqual(quiz, shuffledQuiz)
	
	if return_bool == false {
		t.Errorf("Quiz has not been shuffled")
	}

	_, filename1, _, _ := runtime.Caller(0)
	file1 := path.Join(path.Dir(filename1), "random.csv")
	quiz1 := parseCSVFile(file1)
	shuffledQuiz1 := shuffleQuiz(quiz1)
	return_bool1 := reflect.DeepEqual(quiz1, shuffledQuiz1)

	if return_bool1 == false {
		t.Errorf("Quiz has not been shuffled")
	}

}