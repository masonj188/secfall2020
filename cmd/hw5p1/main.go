package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

var freqChart = map[rune]float32{'a': 8.12, 'b': 1.49, 'c': 2.71, 'd': 4.32, 'e': 12.02, 'f': 2.3, 'g': 2.03, 'h': 5.92, 'i': 7.31, 'j': 0.1, 'k': 0.69, 'l': 3.98, 'm': 2.61, 'n': 6.95, 'o': 7.68, 'p': 1.82, 'q': 0.11, 'r': 6.02, 's': 6.28, 't': 9.1, 'u': 2.88, 'v': 1.11, 'w': 2.09, 'x': 0.17, 'y': 2.11, 'z': 0.07}

func main() {
	// Read the input from stdin, exit if there's an error
	inputBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error reading from stdin")
	}

	// Convert any ascii input to lower case
	inputBytes = []byte(strings.ToLower(string(inputBytes)))

	// Only take the ascii characters from 87 to 122 (all the lowercase letters)
	strBuilder := strings.Builder{}
	for _, c := range inputBytes {
		if c > 96 && c < 123 {
			strBuilder.WriteByte(c)
		}
	}
	input := strBuilder.String()

	inputFreq := map[rune]int{}
	total := len(input)

	// Make sure if a character isn't in the input it's frequency of 0 is recorded
	for i := 96; i < 123; i++ {
		inputFreq[rune(i)] = 0
	}
	// Record each character's frequency
	for _, c := range input {
		inputFreq[c]++
	}

	for k, v := range inputFreq {
		freq := float32(v) / float32(total)
		freqDiff := float64(freq*100 - (freqChart[k]))
		if math.Abs(freqDiff) > 5 { // If the frequency we recorded is 5% greater or lower than the base frequency, print it out
			fmt.Printf("Letter %s differs by %.1f%%.\n", string(k), freqDiff)
		}
	}

}
