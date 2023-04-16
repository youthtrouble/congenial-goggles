package gpt

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var key = "kokoro"
var duration = time.Minute * time.Duration(12)
var chatCompletionCache = cache.New(duration, duration)

func retrieveCachedChatCompletioFormat() ([]chatComletionFormat, bool) {
	if value, ok := getCache(); ok {
		chatCompletions := value.([]chatComletionFormat)
		return reverseArray(chatCompletions), true
	}

	return nil, false
}

func appendNewCacheEntry(role chatRole, text string) {

	if value, ok := getCache(); ok {
		chatCompletions := value.([]chatComletionFormat)
		chatCompletions = append(chatCompletions, chatComletionFormat{role: role, text: text})
		setCache(chatCompletions)
	} else {
		var chatCompletions []chatComletionFormat
		chatCompletion := &chatComletionFormat{
			role: role, 
			text: text,
		}

		chatCompletions = append(chatCompletions, *chatCompletion)
		setCache(chatCompletions)
	}
}

func setCache(value interface{}) {

	chatCompletionCache.Set(key, value, duration)

}

func getCache() (interface{}, bool) {

	return chatCompletionCache.Get(key)

}

//change from lifo to  fifo
func reverseArray(arr []chatComletionFormat) []chatComletionFormat{

	for i, j := 0, len(arr)-1; i<j; i, j = i+1, j-1 {
	   arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

type chatComletionFormat struct {
	role chatRole
	text string
}

type chatRole string

func (role chatRole) String() string {
	return string(role)
}

const (
	assistant chatRole = "assistant"
	user      chatRole = "user"
)
