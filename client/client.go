package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func GetInputData(day int) (string, error) {
	data, err := sendRequest("GET", day, "/input", nil)
	if err != nil {
		return "", fmt.Errorf("failed to get input: %s", err.Error())
	}
	
	return strings.TrimSpace(string(data)), nil
}

func SubmitAnswer(day int, answer int, part int) (string, error) {
	form := url.Values{}
	form.Add("level", strconv.Itoa(part))
	form.Add("answer", strconv.Itoa(answer))

	data, err := sendRequest("POST", day, "/answer", strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to send answer: %s", err.Error())
	}
	return string(data), nil
}

func sendRequest(method string, day int, endpoint string, body io.Reader) (data []byte, err error) {
	url := "https://adventofcode.com/2021/day/" + strconv.Itoa(day) + endpoint
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err.Error())
	}

	request.AddCookie(&http.Cookie{Name: "session", Value: os.Getenv("AOC_SESSION")})
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to get OK response: %s", err.Error())
	}

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %s", err.Error())
	}
	return data, nil
}
