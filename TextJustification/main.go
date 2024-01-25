package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	length, line := 0, []string{}
	res := []string{}

	for i := 0; i < len(words); i++ {
		if length+len(words[i])+len(line) > maxWidth {
			extra_space := maxWidth - length
			spaces := extra_space / max(1, len(line)-1)
			remainder := extra_space % max(1, len(line)-1)

			for j := 0; j < max(1, len(line)-1); j++ {
				line[j] += strings.Repeat(" ", spaces)
				if remainder > 0 {
					line[j] += " "
					remainder--
				}
			}
			res = append(res, strings.Join(line, ""))
			length, line = 0, []string{}
		}

		line = append(line, words[i])
		length += len(words[i])
	}

	last_line := strings.Join(line, " ")
	trail_space := maxWidth - len(last_line)
	res = append(res, last_line+strings.Repeat(" ", trail_space))

	return res
}

func main() {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	fmt.Println(fullJustify(words, 16))
}
