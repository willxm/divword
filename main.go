package main

import (
	"io/ioutil"
	"os"
	"unicode"

	"strings"

	"github.com/huichen/sego"
)

func readFile(path string) []byte {
	fi, err := os.Open(path)
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
	//load dictionary
	segmenter.LoadDictionary("data/dictionary.txt")
	//load news data
	text := readFile("data/news.txt")
	//load stopword
	stopwords := readFile("data/stopwords.txt")
	sws := strings.Split(string(stopwords), "\n")

	segments := segmenter.Segment(text)

	texts := sego.SegmentsToSlice(segments, false)

	fo, err := os.Create("sego.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	l := len(sws)

	for _, v := range texts {
		var i int
		i = 0
		if unicode.IsLetter([]rune(v)[0]) {
			for _, vs := range sws {
				if v != vs {
					i++
				}
			}
		}
		if i == l {
			fo.WriteString(v + "\t")
		}
	}
}
