package files

import (
	"fmt"
	"io"
	"log"
	"os"
)

func BufferedRead() {

	f, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	buf := make([]byte, 8)

	for {
		n, err := f.Read(buf)
		// n, err := io.ReadAtLeast(f, buf, 8)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if err == io.EOF {
			break
		}

		fmt.Print(string(buf[:n]))
	}
}
