package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	var mainErr error
	defer func() {
		if mainErr != nil {
			log.Fatalf("fatal error encountered: %v", mainErr)
		}
	}()

	b, err := ioutil.ReadFile(fmt.Sprintf("year%c2018%cdays%ctwo%cinput.txt", os.PathSeparator, os.PathSeparator, os.PathSeparator, os.PathSeparator))
	if err != nil {
		mainErr = errors.Wrap(err, "read input file")
		return
	}

	input := strings.Split(string(b), "\n")
	var twos, threes int

	for _, s := range input {
		var foundTwo, foundThree bool
		alpha := make(map[rune]int)

		for _, r := range s {
			if _, exists := alpha[r]; !exists {
				alpha[r] = 1
				continue
			}

			alpha[r]++
		}

		for _, v := range alpha {
			if foundTwo && foundThree {
				break
			}

			if v == 2 && !foundTwo {
				foundTwo = true
				twos++
			} else if v == 3 && !foundThree {
				foundThree = true
				threes++
			}
		}
	}

	fmt.Println("checksum:", twos*threes)
}
