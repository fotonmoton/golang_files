package regex

import (
	"fmt"
	"regexp"
)

func DifferentPatternSameResult() {
	text := "can you find third word in this string?"

	number := regexp.MustCompile(`[a-z]{4}`)

	found := number.Find([]byte(text))

	fmt.Printf("the word is: \"%s\" ", string(found))

	same := regexp.MustCompile(`\w{4}`)

	sameNumber := same.Find([]byte(text))

	fmt.Printf("same: \"%s\"\n", string(sameNumber))

}
