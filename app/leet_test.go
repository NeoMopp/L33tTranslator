package main

import (
	"reflect"
	"testing"
)

var testMatrix = [][]string{
	{"a", "z"}, {"b", "y"}, {"c", "x"},
	{"d", "w"}, {"e", "v"}, {"f", "u"},
	{"g", "t"}, {"h", "s"}, {"i", "r"},
	{"j", "q"}, {"k", "p"}, {"l", "o"},
	{"m", "n"}, {"n", "m"}, {"o", "l"},
	{"p", "k"}, {"q", "j"}, {"r", "i"},
	{"s", "h"}, {"t", "g"}, {"u", "f"},
	{"v", "e"}, {"w", "d"}, {"x", "c"},
	{"y", "b"}, {"z", "a"}}

func TestReadCsv(t *testing.T) {
	result := ReadCsv("../test_files/test_matrix.csv")
	if !reflect.DeepEqual(result, testMatrix) {
		t.Fail()
	}
}

func TestCsvRecordsToMap(t *testing.T) {

}

func TestLeetHash(t *testing.T) {

}
