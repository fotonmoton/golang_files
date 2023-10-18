package regex

import (
	"fmt"
	"regexp"
)

var standard = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`

func Email() {
	testEmail := "test-email@some.interesting.domain.bang"

	exactPattern := regexp.MustCompile(`test-email@some.interesting.domain.bang`)

	if exactPattern.Match([]byte(testEmail)) {
		fmt.Println("exact pattern matched")
	}

	anyName := regexp.MustCompile(`.*@some.interesting.domain.bang`)

	if anyName.Match([]byte(testEmail)) {
		fmt.Println("any name matched")

	}

	anyNameAndDomain := regexp.MustCompile(`.*@.*`)

	if anyNameAndDomain.Match([]byte(testEmail)) {
		fmt.Println("any name and domain matched")
	}

	forcedName := regexp.MustCompile(`[\w-\.]+@.*`)

	if forcedName.Match([]byte(testEmail)) {
		fmt.Println("correct name and any domain matched")
	}

	forcedSubdomain := regexp.MustCompile(`[\w-\.]+@([\w-]+\.)`)

	if forcedSubdomain.Match([]byte(testEmail)) {
		fmt.Println("correct name and correct subdomain matched")
	}

	withMinimumErrors := regexp.MustCompile(standard)

	if withMinimumErrors.Match([]byte(testEmail)) {
		fmt.Println("with minimum errors matched")
	}
}
