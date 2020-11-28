package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/scylladb/go-set/strset"
)

func getNixStore() *strset.Set {
	dir := "/nix/store"
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	paths := strset.New()
	for _, file := range files {
		paths.Add(filepath.Join(dir, file.Name()))
	}

	return paths
}

func parseLinear(list string) *strset.Set {
	text := strings.Split(list, "\n")

	linearPaths := strset.New()

	for i := range text {
		path := strings.TrimSpace(text[i])
		if path != "" {
			linearPaths.Add(strings.TrimSpace(text[i]))
		}
	}

	return linearPaths
}

func parseTree(tree string) *strset.Set {
	text := strings.Split(tree, "\n")
	fmt.Println(text)

	treePaths := strset.New()

	for i := range text {
		test := strings.Split(text[i], " ")
		if len(test) == 1 {
			treePaths.Add(test[0])
			continue
		}

		last := strings.TrimSpace(test[len(test)-1])
		secondToLast := strings.TrimSpace(test[len(test)-2])

		if last == "[...]" {
			treePaths.Add(secondToLast[4:len(secondToLast)])
		} else if last != "" {
			treePaths.Add(last[4:len(last)])
		}
	}

	return treePaths
}

func run(cmd *exec.Cmd) string {
	// out, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(fmt.Sprint(err))
	// 	log.Fatal(err)
	// }

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return ""
	}
	return out.String()
}

// func main() {
// 	argLength := len(os.Args[1:])
// 	if argLength != 1 {
// 		fmt.Println("One argument required!")
// 		return
// 	}
// 	dir := os.Args[1]

///////////////////////////
//////// Tree call ////////
///////////////////////////
// command1 := exec.Command("cp", "treescript.sh", dir)
// _, _ = run(command1)

// command2 := exec.Command("./treescript.sh", dir)
// command2.Dir = "/Users/josephkim/Documents/Senior2020/deplorable"
// out2, _ := run(command2)
// paths := parseTree(out2)

// command3 := exec.Command("rm", "treescript.sh")
// command3.Dir = dir
// _, _ = run(command3)

///////////////////////////
//////// Linear call //////
///////////////////////////
// command1 := exec.Command("cp", "linearscript.sh", dir)
// _, _ = run(command1)

// command2 := exec.Command("./linearscript.sh")
// command2.Dir = dir
// out2, _ := run(command2)
// paths := parseLinear(out2)

// command3 := exec.Command("rm", "linearscript.sh")
// command3.Dir = dir
// _, _ = run(command3)

// nixStore := getNixStore()
// builtDerivations := strset.Intersection(paths, nixStore)
// needToBeBuilt := strset.Difference(paths, nixStore)

// fmt.Println(paths.Size())
// fmt.Println(nixStore.Size())
// fmt.Println(builtDerivations.Size())
// fmt.Println(needToBeBuilt.Size())
// fmt.Println(needToBeBuilt)

// 	command := exec.Command("/bin/bash", "-c", "nix show-derivation /nix/store/yfk28ll9iaf4k8k4nic3fhr6jadhxbkn-rust_deplorable-0.1.0.drv", dir)
// 	// command := exec.Command("/bin/bash", "-c", "echo hi && echo bye", dir)
// 	out, _ := run(command)
// 	fmt.Println()
// 	fmt.Println()

// 	fmt.Printf("%s", out)
// }
