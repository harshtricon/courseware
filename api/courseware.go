package api

import (
	"encoding/json"
	"fmt"
)

const API_URL = "https://dle-cms-assemblyapi.mheducation.com/v3/container/%s/%s"

func FetchCourseware(containerID, publishID, jwtToken string) (string, error) {
	url := fmt.Sprintf(API_URL, containerID, publishID)
	response, err := MakeRequest(url, jwtToken)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		return "", err
	}

	formattedResult, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", err
	}

	return string(formattedResult), nil
}
