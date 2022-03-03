package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string, repeat int) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		repeat_line := line
		for i := 1; i < repeat; i++ {
			repeat_line += "-" + line
		}
		fmt.Fprintln(w, repeat_line)
	}
	return w.Flush()
}

func main() {
	var file_in, file_out string
	var repeat int

	// flags declaration using flag package
	flag.StringVar(&file_in, "i", "test.txt", "name of input file, default is test.txt")
	flag.StringVar(&file_out, "o", "test.out.txt", "name of output file, default is test.out.txt")
	flag.IntVar(&repeat, "n", 0, "number of repetitions, default is 1")
	flag.Parse()

	lines, err := readLines(file_in)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for i, line := range lines {
		fmt.Println(i, line)
	}

	if err := writeLines(lines, file_out, repeat); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
}
