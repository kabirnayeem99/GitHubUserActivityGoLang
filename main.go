package main

import (
	"fmt"
	"github.com/kabirnayeem99/GitHubUserActivityGoLang/internal/githubactivity"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalln("Usage: github-activity [username]")
		os.Exit(1)
	}

	username := os.Args[1]
	responseEvents, err := githubactivity.GetGithubUserEvents(username)

	if err != nil {
		log.Fatalln("Failed to call API, because, ", err)
		os.Exit(1)
	}

	fmt.Println("Output:")
	for _, re := range responseEvents {
		switch re.Type {
		case string(githubactivity.CreateEvent):
			fmt.Printf(" - [%v] Created %v %v at %v\n", re.CreatedAt, re.Payload.Ref, re.Payload.RefType, re.Repo.Name)
		case string(githubactivity.PushEvent):
			fmt.Printf(" - [%v] Pushed new changes to %v\n", re.CreatedAt, re.Repo.Name)
		case string(githubactivity.WatchEvent):
			fmt.Printf(" - [%v] Started wathing %v\n", re.CreatedAt, re.Repo.Name)
		}
	}
}
