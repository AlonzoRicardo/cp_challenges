package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// To read from input
var reader = bufio.NewReader(os.Stdin)

type Hashdata struct {
	foundHash int
	files     map[string]int
}

// XOR hash
func hash(b []uint8) uint8 { // O(m)
	var hash uint8
	for _, e := range b {
		hash = hash ^ e
	}
	return hash
}

func main() {
	for {
		// Read the frist line to get the number of test cases
		input, _ := reader.ReadString('\n')
		cases, _ := strconv.Atoi(strings.TrimSpace(input))

		// There are no more test cases 0 is read from the input
		if cases == 0 {
			break
		}

		var m = map[uint8]*Hashdata{}
		var totalCols int
		var uniques int

		for i := 0; i < cases; i++ { // O(n)
			// read files and trim space
			str, _ := reader.ReadString('\n')
			file := strings.TrimRight(str, " ")
			// get file as an array of bytes
			fileBytes := []byte(file)

			// get hash from byte array
			var hash uint8 = hash(fileBytes)

			// if the hash is in our map we add 1 to number of coalitions
			if _, ok := m[hash]; ok {
				// if the file is duplicated we count it
				// else add one to uniques
				if _, ok := m[hash].files[file]; ok {
					m[hash].files[file]++
				} else {
					m[hash].files[file] = 1
					uniques++
				}
				// add 1 to coalitions
				m[hash].foundHash++
			} else {
				// when no hash is found we initialize it in our map
				data := &Hashdata{files: make(map[string]int)}
				data.files[file] = 1
				data.foundHash++
				m[hash] = data
				uniques++
			}
			// Add to total number of coalitions
			// the amount of coalitions found for that hash minus the amount of times
			// the actual file was found.
			totalCols += (m[hash].foundHash) - m[hash].files[file]
		}
		fmt.Println(uniques, totalCols)
	}
}

// Complexity O(n*m)
