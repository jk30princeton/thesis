package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// content, err := ioutil.ReadFile("/Users/josephkim/Downloads/names.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// text := string(content)
	// split := strings.Split(text, "\n")

	// for i, s := range split {
	// 	if s == "" {
	// 		continue
	// 	}
	// 	fmt.Println(i, s)
	// }

	command2 := exec.Command("/bin/bash", "-c", "nix-instantiate '<nixpkgs>' -A firefox")
	command2.Dir = dir
	out2 := strings.TrimSpace(run(command2))
	fmt.Println(out2)

	// set1 := strset.New()
	// set1.Add("Hello")

	// set2 := strset.New()
	// set2.Add("Hello")
	// set2.Add("World")

	// set3 := strset.Union(set1, set2)
	// set3.Add("Hello")
	// fmt.Println(set3)
	// fmt.Println(set3.Has("Hello"))
}