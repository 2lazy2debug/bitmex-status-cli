package utils

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func readFileLines(filepath string) (keys map[string]string) {
	f, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return
		}
		parts := strings.Split(line, ":") // GET the line string
		keys[parts[0]] = parts[1]
	}
	return
}
