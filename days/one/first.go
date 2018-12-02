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

	for _, v := range input {
		i, err := strconv.Atoi(v)
		if err != nil {
			mainErr = errors.Wrap(err, "convert int")
			return
		}

		frequency += i
	}

	fmt.Println("resulting frequency:", frequency)
}
