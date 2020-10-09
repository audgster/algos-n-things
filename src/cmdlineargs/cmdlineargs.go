package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for _, filename := range os.Args[1:] {
			file, err := os.Open(filename)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Something went wrong: %s", err)
				continue
			}

			streamFiles(file, filename)
		}
	} else {
		streamFiles(os.Stdin, "Stdin")
	}

}

func streamFiles(file *os.File, filename string) {
	fmt.Printf("Streaming file %s\n", filename)
	input := bufio.NewScanner(file)
	for input.Scan() {
		fmt.Println(input.Text())
	}
}
