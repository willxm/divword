package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode"

	"github.com/huichen/sego"
)

func readFile(path string) []byte {
	fi, err := os.Open("news.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}
	return fd
}

func main() {
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("/Users/willxm/go/src/github.com/huichen/sego/data/dictionary.txt")

	text := readFile("news.txt")
	segments := segmenter.Segment(text)

	texts := sego.SegmentsToSlice(segments, false)

	for _, v := range texts {
		if !unicode.IsLetter([]rune(v)[0]) {

		} else {
			fmt.Println(v)
		}
	}
}
