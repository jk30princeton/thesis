package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/scylladb/go-set/strset"
)

func main() {
	// Can set to any directory that has a file called "default.nix"
	dir := "/Users/josephkim/Documents/Senior2020/deplorable"

	// Tree call
	command1 := exec.Command("./treescript.sh", dir)
	out1, _ := run(command1)
	paths := parseTree(out1)

	nixStore := getNixStore()

	fmt.Println(nixStore)

	// Linear call
	command2 := exec.Command("./linearscript.sh", dir)
	out2, _ := run(command2)
	paths2 := parseLinear(out2)

	fmt.Println()
	fmt.Println(paths2.Size())
	fmt.Println(paths.Size())
}

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

func run(cmd *exec.Cmd) (string, error) {
	var output bytes.Buffer
	var waitGroup sync.WaitGroup

	stdout, _ := cmd.StdoutPipe()
	writer := io.MultiWriter(os.Stdout, &output)

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		io.Copy(writer, stdout)
	}()

	cmd.Run()
	waitGroup.Wait()
	return output.String(), nil
}
