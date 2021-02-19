package utils

import (
	"errors"
	"net/http"
	"regexp"
)

// VerifyTokenUsagePlan ...
func VerifyTokenUsagePlan(huggingfaceToken *string) error {

	url := "https://huggingface.co/api/whoami-v2"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+*huggingfaceToken)

	client := &http.Client{}
	resp, err := client.Do(req)

	if resp.StatusCode == 200 {
		return nil
	}
	return errors.New("token invalide")

}

/*
	|						. ---------------------------- Match exact the string
	|						|			    . ------------ Match any Char for a-zA-Z
	|						| 			   |     . ------- Match 34 times or length of 34
	|						| 			   |     |   . --- Must match exact 34 times */
var r, _ = regexp.Compile("(api|api_org)_[a-zA-Z]{34}$")

// VerifyTokenStringSchema ...
func VerifyTokenStringSchema(token *string) error {
	if !r.MatchString(*token) {
		return errors.New("API-Token doesn`t match schema in form of /(api|api_org)_[a-zA-Z]{34}$")
	}
	return nil
}
