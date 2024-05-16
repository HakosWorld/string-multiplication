package reload

import (
	"log"
	"strconv"
	"strings"
)

func ProcessPrev(words []string, index int, task string) {
	// Find the number after "(XXX,"
	if index+1 < len(words) {
		numStr := strings.TrimRight(words[index+1], ",)")
		num, err := strconv.Atoi(numStr)
		// if 'num' exeeds the previews words, set the 'num' to equal the index

		if index < num {
			log.Fatal("Words in a function is out of range")
		}

		if err == nil {
			// change the previous 'num' words
			for j := index; j >= index-num; j-- {

				if task == "up" {
					words[j] = strings.ToUpper(words[j])
				}
				if task == "low" {
					words[j] = strings.ToLower(words[j])
				}
				if task == "cap" {
					word := words[j]
					//words[j] = strings.Title(strings.ToLower(words[j]))
					if word[0] == '\'' {
						words[j] = "'" + strings.ToUpper(string(word[1])) + strings.ToLower(string(word[2:]))
					} else {
						words[j] = strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
					}
				}

			}
		}
	}
}
