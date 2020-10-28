package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"gopkg.in/yaml.v2"
)

var secretMap map[string]string
var indexMap map[string]int
var repos Repos

type Repos struct {
	Repositories []Repo `yaml:"repos"`
}

type Repo map[string]string

func parseYaml(yamlFile []uint8) Repos {
	var repos Repos
	err := yaml.Unmarshal(yamlFile, &repos)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	return repos
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	pathName := r.URL.Path[1:]
	payload, err := github.ValidatePayload(r, []byte(secretMap[pathName]))
	if err != nil {
		log.Printf("Error validating request body: err=%s\n", err)
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("Could not parse webhook: err=%s\n", err)
		return
	}
	log.Printf("%v\n", event)
	log.Printf("%v\n", github.WebHookType(r))
	switch e := event.(type) {
	// Commit push event
	case *github.PushEvent:
		log.Printf("Push event received: %v", e)
	// Ignore all other events
	default:
		log.Printf("Not a push event. Event type is %s\n", e)
		return
	}
}

func handleRepoRequest(repo string) int {
	/*
	   app := "nix"
	   arg0 := "build"
	   arg1 := "--out-link"
	   arg2 := repos.Repositories[repo]["out"]
	   arg3 := "-f"
	   arg4 := tarball

	   cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4)
	   stdout, err := Command.Output()

	   if err != nil {
	       log.Printf("Failed to execute %s", cmd)
	   }*/
	log.Printf("Handling build")
	return 0
}

func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "", "YAML file to parse.")
	flag.Parse()

	if fileName == "" {
		log.Println("Please provide YAML file by using -f option")
		return
	}

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("Error reading YAML: Err is %s\n", err)
		return
	}

	repos := parseYaml(yamlFile)

	//log.Println("server started")
	secretMap = make(map[string]string, len(repos.Repositories))

	for i := 0; i < len(repos.Repositories); i++ {
		log.Printf(repos.Repositories[i]["name"])
		http.HandleFunc("/"+repos.Repositories[i]["name"], handleWebhook)
		secretMap[repos.Repositories[i]["name"]] = repos.Repositories[i]["secret"]
		indexMap[repos.Repositories[i]["name"]] = i
	}
	log.Fatal(http.ListenAndServe(":8088", nil))
}
