package regex

import (
	"fmt"
	"regexp"

	"github.com/fotonmoton/golang_files/check"
)

func SimpleIndex() {
	text := "how to regex?"

	regex, err := regexp.Compile(`.....\?`)

	check.Check(err)

	found := regex.FindIndex([]byte(text))

	if found == nil {
		fmt.Println("can't find")
		return
	}

	fmt.Println("Found indexes:", found, "string:", text[found[0]:found[1]])
}
