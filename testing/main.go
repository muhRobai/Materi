package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	resp, _ := getValue("eat")
	log.Println(resp)
}

func getValue(coder string) (string, error) {
	isRealCode := false
	log.Println(coder)
	var codeType []string
	if coder == "eat" {
		isRealCode = true
		codeType = append(codeType, "eat")
	}

	if coder == "code" {
		isRealCode = true
		codeType = append(codeType, "code")
	}

	if coder == "sleep" {
		isRealCode = true
		codeType = append(codeType, "sleep")
	}

	if !isRealCode {
		return "", errors.New("you are not coder")
	}

	var result string
	for _, p := range codeType {
		result += fmt.Sprintf("%s", p)
	}

	return result, nil

}
