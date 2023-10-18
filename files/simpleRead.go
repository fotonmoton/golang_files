package files

import (
	"fmt"
	"os"

	"github.com/fotonmoton/golang_files/check"
)

func SimpleRead() {
	content, err := os.ReadFile("main.go")

	check.Check(err)

	fmt.Println(string(content))
}
