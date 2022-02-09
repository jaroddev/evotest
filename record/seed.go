package record

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Here we read a file
func readFile(filename string) os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return *file
}

func splitLines(file os.File) []string {
	buffer := bufio.NewScanner(&file)

	lines := make([]string, 0)

	for buffer.Scan() { // internally, it advances token based on sperator
		lines = append(lines, buffer.Text())
	}

	return lines
}

func Read(filename string) []string {
	file := readFile(filename)
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	return splitLines(file)
}

func ConvertLinesToSeeds(lines []string) []int64 {
	seeds := make([]int64, 0)

	for _, line := range lines {
		seed, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, seed)
	}

	return seeds
}
