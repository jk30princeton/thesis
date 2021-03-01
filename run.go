package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/scylladb/go-set/strset"
)

func main() {
	content, err := ioutil.ReadFile("/Users/josephkim/Downloads/names.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	split := strings.Split(text, "\n")

	for i, s := range split {
		if s == "" {
			continue
		}
		fmt.Println(i, s)
	}
	set1 := strset.New()
	set1.Add("Hello")

	set2 := strset.New()
	set2.Add("Hello")
	set2.Add("World")

	set3 := strset.Union(set1, set2)
	set3.Add("Hello")
	fmt.Println(set3)
	fmt.Println(set3.Has("Hello"))
}
