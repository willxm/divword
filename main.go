package main

import (
	"io/ioutil"
	"log"
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
	log.Printf("载入文本")
	text := readFile("data/news.txt")

	//load stopword
	log.Printf("载入停用词词典")
	stopwords := readFile("data/stopwords.txt")
	sws := strings.Split(string(stopwords), "\n")

	//process segment
	log.Printf("分词...")
	segments := segmenter.Segment(text)
	texts := sego.SegmentsToSlice(segments, false)

	fo, err := os.Create("sego.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	l := len(sws)
	//process stopwords
	log.Printf("去除停用词...")
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
	log.Printf("完成")
}
