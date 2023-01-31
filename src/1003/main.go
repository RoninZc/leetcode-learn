package main

import "fmt"

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
}

func commonChars(words []string) []string {
	dics := make([][26]int, len(words))
	for i, word := range words {
		for _, b := range []byte(word) {
			dics[i][b-'a']++
		}
	}
	data := make([]string, 0)
	for i := 0; i < 26; i++ {
		if counter := dics[0][i]; counter != 0 {
			for j := 1; j < len(dics); j++ {
				if tmp := dics[j][i]; tmp > 0 {
					if tmp < counter {
						counter = tmp
					}
				} else {
					counter = -1
					break
				}
			}
			for j := 0; j < counter; j++ {
				data = append(data, string(byte(i+'a')))
			}
		}
	}

	return data
}
