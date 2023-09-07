package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	apiURL = "https://api.openai.com/v1/chat/completions"
	apiKey = ""
)

func readFile(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func generateGPTResponse(data string) (string, error) {
	requestData := map[string]interface{}{
		"model":       "gpt-3.5-turbo",
		"messages":    []map[string]interface{}{{"role": "user", "content": data}},
		"temperature": 0.7,
	}

	requestDataJSON, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestDataJSON))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return "API Response:" + string(responseBody), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	inputFileName := os.Args[1]
	data, err := readFile(inputFileName)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	resp, err := generateGPTResponse(data)
	if err != nil {
		fmt.Printf("Error processing data: %v\n", err)
		return
	}

	_ = resp
	// fmt.Println(data)
	// fmt.Println(resp)
}
