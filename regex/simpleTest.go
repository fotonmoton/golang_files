package regex

import (
	"fmt"
	"regexp"

	"github.com/fotonmoton/golang_files/check"
)

func SimpleTest() {
	text := "how to regex?"

	found, err := regexp.MatchString("regex", text)

	check.Check(err)

	fmt.Println("Is there a match?", found)
}
