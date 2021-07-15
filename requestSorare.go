package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var SorareUrl = "https://api.sorare.com/graphql"

func MarshalQuery(qS string, vars string, op string) ([]byte, error) {

	var gqlr GraphQLRequest
	gqlr.Query = qS
	gqlr.Variables = vars
	gqlr.OperationName = op
	mq, err := json.Marshal(gqlr)
	if err != nil {
		return []byte{}, err
	}
	return mq, nil
}

func MarshalQuery2(qS string, vars string, op string) (bytes.Buffer, error) {
	var mq bytes.Buffer
	var gqlr GraphQLRequest

	gqlr.Query = qS
	gqlr.Variables = vars
	gqlr.OperationName = op
	if err := json.NewEncoder(&mq).Encode(gqlr); err != nil {
		return bytes.Buffer{}, err
	}

	return mq, nil
}

func CRREquest() (string, error) {

	mq, err := MarshalQuery(CR, VCR, "")
	if err != nil {
		return "", err
	}
	//fmt.Println("QUERYTOTAL", string(mq))
	request, err := http.NewRequest("POST", SorareUrl, strings.NewReader(string(mq)))
	//fmt.Println("REQUEST:", request)
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 10}
	rsp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	return string(data), err
}

func LoginRequest() (string, error) {

	mq, err := MarshalQuery(SORARE, VS, OP_SIGNIN_MUTATION)
	if err != nil {
		return "", err
	}
	//fmt.Println("QUERYTOTAL", string(mq))
	request, err := http.NewRequest("POST", SorareUrl, strings.NewReader(string(mq)))
	fmt.Println("REQUEST:", string(mq))
	if err != nil {
		return "", err
	}

	_, csrf, cookie := GetSalt()
	request.AddCookie(cookie)
	request.Header.Add("X-CSRF-Token", csrf)
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	rsp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	return string(data), err

}
