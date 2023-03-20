package gpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/youthtrouble/congenial-goggles/utils"
)

const openAIURL = "https://api.openai.com/v1"

func RetrieveOpenAICompletions(prompt string) (*string, error) {

	request := completionsRequest{
		//might want to change the model here to something stronger
		//glenn 50 might do lmao
		Model:  "text-davinci-003",
		Prompt: prompt,
		//might also want to change this to something lowe
		//open ai charges per 1000 tokens
		MaxTokens:   500,
		//please check the openAI docs for a better overview of the temperature
		//parameter controls.
		//--but for Nigerians, -- the blood fit dey hhpt so it'll do crazy things
		//i.e take more risks and not sound bland
		Temperature: 1,
	}

	var response completionsResponse
	err := executeOpenAIRequest("POST", "completions", request, &response)
	if err != nil {
		return nil, err
	}

	return &response.Choices[0].Text, nil
}

func executeOpenAIRequest(method, endpoint string, requestData, destination interface{}) error {

	url := fmt.Sprintf("%s/%s", openAIURL, endpoint)
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	var req *http.Request

	if requestData == nil {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return err
		}
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(requestBody))
		if err != nil {
			return err
		}
	}

	req.Header.Set("Authorization", "Bearer "+utils.UseEnvOrDefault("OPENAI_API_KEY", "sk-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
	req.Header.Set("Content-Type", "application/json")

	var response *http.Response
	response, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	responseCode := response.StatusCode
	if responseCode != 200 && responseCode != 201 {
		return errors.New("error processing request")
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(responseBody, destination)
}
