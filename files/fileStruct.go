package files

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/fotonmoton/golang_files/check"
)

func FileStructRead() {
	file, err := os.Open("main.go")
	check.Check(err)

	first5bytes := make([]byte, 7)
	n1, err := file.Read(first5bytes)
	check.Check(err)

	fmt.Printf("%d bytes: %s\n", n1, string(first5bytes[:n1]))

	o3, err := file.Seek(2, 0)
	check.Check(err)

	b2 := make([]byte, 2)
	n3, err := io.ReadAtLeast(file, b2, 2)
	check.Check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b2))

	r4 := bufio.NewReader(file)
	b4, err := r4.Peek(5)
	check.Check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// _, err = file.Seek(0, 0)
	// check.Check(err)

	// r4 = bufio.NewReader(file)

	b4, err = r4.Peek(20)
	check.Check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	file.Close()
}
