package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var mainErr error
	defer func(){
		if mainErr != nil {
			log.Fatalf("fatal error encountered: %v", mainErr)
		}
	}()

	b, err := ioutil.ReadFile(fmt.Sprintf("days%cone%cinput.txt", os.PathSeparator, os.PathSeparator))
	if err != nil {
		mainErr = errors.Wrap(err, "read input file")
		return
	}

	input := strings.Split(string(b), "\n")
	var frequency int
	var twice *int
	frequencies := []int{0}

	for twice == nil {
		for _, v := range input {
			i, err := strconv.Atoi(v)
			if err != nil {
				mainErr = errors.Wrap(err, "convert int")
				return
			}

			frequency += i

			for _, fs := range frequencies {
				if fs == frequency {
					twice = &frequency
					break
				}
			}

			if twice != nil {
				break
			}

			frequencies = append(frequencies, frequency)
		}
	}

	fmt.Println("twice:", *twice)
}
