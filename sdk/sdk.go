package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/maxhaensel/test-go-sdk/utils"
)

type input struct {
	Inputs string `json:"inputs"`
}

// InferencePredictor ...
type InferencePredictor struct {
	HuggingfaceToken *string
	Model            *string
}

// Predict ..
func (i InferencePredictor) Predict(predict string) string {
	url := strings.Join([]string{"https://api-inference.huggingface.co", "models", *i.Model}, "/")

	jsonStr, err := json.Marshal(input{
		Inputs: predict,
	})

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+*i.HuggingfaceToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

/*
	|						. ---------------------------- Match exact the string
	|						|			    . ------------ Match any Char for a-zA-Z
	|						| 			   |     . ------- Match 34 times or length of 34
	|						| 			   |     |   . --- Must match exact 34 times */
var r, _ = regexp.Compile("(api|api_org)_[a-zA-Z]{34}$")

// Predict ..
func (i InferencePredictor) SetToken(token *string) error {
	if err := utils.VerifyTokenStringSchema(token); err != nil {
		return err
	}

	if err := utils.VerifyTokenUsagePlan(token); err != nil {
		return err
	}
	i.HuggingfaceToken = token
	return nil
}

// ValidateToken ..
func (i InferencePredictor) ValidateToken() error {
	if i.HuggingfaceToken == nil {
		return errors.New("not token was set!, set Token or use SetToken method")
	}
	if err := utils.VerifyTokenStringSchema(i.HuggingfaceToken); err != nil {
		return err
	}

	if err := utils.VerifyTokenUsagePlan(i.HuggingfaceToken); err != nil {
		return err
	}
	return nil
}
