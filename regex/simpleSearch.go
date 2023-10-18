package regex

import (
	"fmt"
	"regexp"

	"github.com/fotonmoton/golang_files/check"
)

func SimpleSearch() {
	text := "how to regex?"

	regex, err := regexp.Compile("regex")

	check.Check(err)

	found := regex.Find([]byte(text))

	fmt.Println("we got him:", string(found))
}
