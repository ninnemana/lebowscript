package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	parse("DUDE")
}

func parse(name string) {
	file, err := os.Open("script.txt")
	if err != nil {
		panic(err)
	}

	lines := make(map[string][]string, 0)

	scanner := bufio.NewScanner(file)
	var lastCharacter string
	var line string
	for scanner.Scan() {
		trimmed := strings.TrimSpace(scanner.Text())
		// empty line check
		if trimmed == "" {
			continue
		}

		last := trimmed[len(trimmed)-1:]
		// annoying scene message
		if last == ":" {
			continue
		}

		// weird hack for inconsistency
		if trimmed == "THE DUDE" {
			trimmed = "DUDE"
		}

		// seems if it's uppercase, it's a character name callout
		if strings.ToUpper(trimmed) == trimmed {
			lines[lastCharacter] = append(lines[lastCharacter], line)
			line = ""
			lastCharacter = trimmed
			if _, ok := lines[trimmed]; !ok {
				lines[trimmed] = make([]string, 0)
				continue
			}
		} else {
			line = fmt.Sprintf("%s %s", line, strings.Replace(trimmed, "\n", " ", 1))
		}
	}

	for char, lines := range lines {

		if strings.ToLower(char) == strings.ToLower(name) {
			for _, line := range lines {
				fmt.Println(line)
			}
		}
	}
}
