package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/scylladb/go-set/strset"
)

func main() {
	// Tree
	content, err := ioutil.ReadFile("snsDependencies.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := strings.Split(string(content), "\n")

	tree_paths := strset.New("")

	counter := make(map[string]int)
	for i := range text {
		test := strings.Split(text[i], " ")
		if strings.TrimSpace(test[len(test)-1]) == "[...]" {
			path := strings.TrimSpace(test[len(test)-2])
			//fmt.Println(path[4:len(path)])
			tree_paths.Add(path[4:len(path)])
			counter[path[4:len(path)]]++
		} else if strings.TrimSpace(test[len(test)-1]) != "" {
			path := strings.TrimSpace(test[len(test)-1])
			//fmt.Println(path[4:len(path)])
			tree_paths.Add(path[4:len(path)])
			counter[path[4:len(path)]]++
		}
	}

	// Linear
	content, err = ioutil.ReadFile("snsDependencies2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text = strings.Split(string(content), "\n")

	linear_paths := strset.New("")

	for i := range text {
		linear_paths.Add(strings.TrimSpace(text[i]))
		//fmt.Println(strings.TrimSpace(text[i]))
	}

	s3 := strset.Intersection(tree_paths, linear_paths)
	fmt.Println("Intersection")
	fmt.Println(s3.Size())
	fmt.Println(s3)

	s4 := strset.Difference(tree_paths, linear_paths)
	fmt.Println("Difference of tree_paths and linear_paths")
	fmt.Println(s4.Size())
	fmt.Println(s4)

	s5 := strset.Difference(linear_paths, tree_paths)
	fmt.Println("Difference of linear_paths and tree_paths")
	fmt.Println(s5.Size())
	fmt.Println(s5)

	fmt.Println(counter)
	fmt.Println(len(counter))
	fmt.Println(linear_paths.Size())
}
