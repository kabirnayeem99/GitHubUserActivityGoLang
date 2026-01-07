package githubactivity

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func GetGithubUserEvents(username string) (GitHubEventsApiResponse, error) {
	if len(username) == 0 {
		return GitHubEventsApiResponse{}, errors.New("Empty user name.")
	}
	response, err := http.Get("https://api.github.com/users/" + username + "/events")

	if err != nil {
		log.Println("Failed to call API, because, ", err)
		return GitHubEventsApiResponse{}, err
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Println("Failed to read API response, because, ", err)
		return GitHubEventsApiResponse{}, err
	}

	var apiResponse GitHubEventsApiResponse
	json.Unmarshal(responseData, &apiResponse)

	return apiResponse, nil
}
