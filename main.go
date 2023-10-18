package main

import (
	"fmt"

	"github.com/fotonmoton/golang_files/files"
	"github.com/fotonmoton/golang_files/regex"
)

var examples = []struct {
	name string
	fun  func()
}{
	{"SimpleTest", regex.SimpleTest},
	{"SimpleSearch", regex.SimpleSearch},
	{"SimpleIndex", regex.SimpleIndex},
	{"DifferentPatternSameResult", regex.DifferentPatternSameResult},
	{"FindAndReplace", regex.FindAndReplace},
	{"Email", regex.Email},
	{"SimpleRead", files.SimpleRead},
	{"FileStructRead", files.FileStructRead},
	{"FileScanner", files.FileScanner},
	{"FileScannerWords", files.FileScannerWords},
	{"BufferedRead", files.BufferedRead},
}

func main() {
	for _, example := range examples {
		fmt.Print(example.name, ": ")
		fmt.Println()
		example.fun()
	}
}
