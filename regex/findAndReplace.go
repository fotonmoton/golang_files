package regex

import (
	"fmt"
	"regexp"
)

func FindAndReplace() {
	rt := regexp.MustCompile(`r.t`)
	fmt.Println(rt.ReplaceAllString("rat cat rot dog", "ram"))
	// prints "ram cat ram dog"

	dashed := regexp.MustCompile(`-.*-`)
	fmt.Println(dashed.ReplaceAllString("-rasjdkajnsdt-hello world", ""))
	// prints "hello world"
}
