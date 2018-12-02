package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var mainErr error
	defer func(){
		if mainErr != nil {
			log.Fatalf("fatal error encountered: %v", mainErr)
		}
	}()

	b, err := ioutil.ReadFile(fmt.Sprintf("days%ctwo%cinput.txt", os.PathSeparator, os.PathSeparator))
	if err != nil {
		mainErr = errors.Wrap(err, "read input file")
		return
	}

	input := strings.Split(string(b), "\n")
	var protoFabric []string

	for _, s := range input {
		for _, innerS := range input {
			if s == innerS {
				continue
			}

			var diff int

			for k, r := range s {
				if r != rune(innerS[k]) {
					diff++
				}
			}

			if diff == 1 {
				protoFabric = append(protoFabric, s, innerS)
				break
			}
		}

		if len(protoFabric) != 0 {
			break
		}
	}

	var i int
	fmt.Println("common letters:", strings.Map(func(r rune) rune{
		defer func(){ i++ }()

		if r != rune(protoFabric[1][i]) {
			return -1
		}

		return r
	}, protoFabric[0]))
}
