package main

import (
	"testing"
	"runtime"
	"path"
	"reflect"
)

func TestShuffle(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	file := path.Join(path.Dir(filename), "problems.csv")
	var quiz []Quiz
	quiz = parseCSVFile(file)
	shuffledQuiz := shuffleArray(quiz)
	return_bool := reflect.DeepEqual(quiz, shuffledQuiz)
	if return_bool == false {
		t.Errorf("Quiz has not been shuffled")
	}
}