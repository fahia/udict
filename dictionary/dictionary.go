package dictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	ResultType  string       `json:"result_type"`
	Definitions []Definition `json:"list"`
}

type Definition struct {
	Definition string `json:"definition"`
	ThumbsUp   int    `json:"thumbs_up"`
	Word       string `json:"word"`
	Example    string `json:"example"`
}

const UrbanDictionaryAPI = "https://api.urbandictionary.com/v0/define?term="

func Define(word string) Response {
	definedObject := fmt.Sprintf("%s%v", UrbanDictionaryAPI, word)
	res, err := http.Get(definedObject)
	responseBody, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var response Response
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return response
}
