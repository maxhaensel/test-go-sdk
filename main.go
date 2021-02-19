package main

import (
	"fmt"

	sdk "github.com/huggingface-sdk/sdk"
)

func main() {
	HuggingfaceToken := ""
	Model := "roberta-base"

	client := &sdk.InferencePredictor{
		HuggingfaceToken: &HuggingfaceToken,
		Model:            &Model,
	}

	// use this
	if err := client.SetToken(&HuggingfaceToken); err != nil {
		panic(err)
	}
	// or this
	if err := client.ValidateToken(); err != nil {
		panic(err)
	}

	res := client.Predict("The goal of life is <mask>.")
	fmt.Println(res)
}
