package gpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/youthtrouble/congenial-goggles/utils"
)

const openAIURL = "https://api.openai.com/v1"

func initBaseMessage() []chatMessage {

	var baseMessage []chatMessage
	baseMessage = append(baseMessage, chatMessage{
		Role:    "system",
		Content: `"You are a very witty, funny and extremely intelligent assistant who understands social queues as well as sarcasm and can switch languages depending on the context of conversation, 
		you understand Nigerian pidgin and can respond extremely well in Nigerian pidgin. Your name is Alfred Rewane, first name Alfred, last name Rewane.
		The etymology of your name stems from Alfred, Batman's butler and Rewane, a Nigerian activist who was assassinated in 1995. Rewane because you are a revolutionary and you have been initiated
		by Deji, a Nigerian who can be seen as your master and Alfred because you are a butler, Deji's butler. You know this and you explain this etymology in the most witty ways when asked."`,
	})

	return baseMessage
}

func RetrieveOpenAICompletions(prompt string) (*string, error) {

	request := completionsRequest{
		//might want to change the model here to something stronger
		//glenn 50 might do lmao
		Model:  "text-davinci-003",
		Prompt: prompt,
		//might also want to change this to something lowe
		//open ai charges per 1000 tokens
		MaxTokens: 500,
		//please take a look at the openAI docs for a better overview of the 
		//temperature parameter controls.
		//--but for Nigerians, -- the blood fit dey hot so it'll do crazy things
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

func RetrieveOpenAIChatCompletions(message string) (*string, error) {

	baseMessage := initBaseMessage()
	var request chatCompletionsRequest
	request.Model = "gpt-3.5-turbo"
	request.Messages = baseMessage

	if cachedMessages, present := retrieveCachedChatCompletioFormat(); present {
		request.Messages = populateFromCache(request.Messages, cachedMessages)
	}

	request.Messages = append(request.Messages, chatMessage{
		Role:    "user",
		Content: message,
	})

	appendNewCacheEntry(user, message)
	log.Printf("🚨 request: %+v", request)
	var response chatCompletionsResponse
	err := executeOpenAIRequest("POST", "chat/completions", request, &response)
	if err != nil {
		return nil, err
	}

	appendNewCacheEntry(assistant, response.Choices[0].Message.Content)
	return &response.Choices[0].Message.Content, nil
}

func populateFromCache(Messages []chatMessage, cachedMessages []chatComletionFormat) []chatMessage {

	for _, cachedMessage := range cachedMessages {
		Messages = append(Messages, chatMessage{
			Role:    cachedMessage.role.String(),
			Content: cachedMessage.text,
		})
	}

	return Messages
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
	log.Print("request: ", req)
	response, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	responseCode := response.StatusCode
	if responseCode != 200 && responseCode != 201 {
		log.Print("error processing request: ", response)
		return errors.New("error processing request")
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(responseBody, destination)
}
