package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jameskeane/bcrypt"
	"github.com/machinebox/graphql"
)

// make a get and take a salt, a csrf and a cookie
func getSalt() ([]byte, string, *http.Cookie) {

	var cookie *http.Cookie
	buser := os.Getenv("BACKEND_USER")
	url := BASE_URL + "/api/v1/users/" + buser
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	csrf := resp.Header.Get("CSRF-TOKEN")
	cookies := resp.Cookies()
	for _, c := range cookies {
		if c.Name == "_sorare_session_id" {
			cookie = c
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		return bodyBytes, csrf, cookie

	} else {
		return []byte{}, "", &http.Cookie{}
	}
}

// retuns s of type Salt
func readSalt(salt []byte) (s Salt) {

	err := json.Unmarshal(salt, &s)

	if err != nil {

		fmt.Printf("Error decodificando: %v\n", err)
	}
	return s
}

//Create a payload with email and pass hashed info
func setPayload() (payload TUser) {

	salt, _, _ := getSalt() //salt in bytes
	s := readSalt(salt)     //salt in struct

	payload.Email = os.Getenv("BACKEND_USER")
	payload.Password = os.Getenv("BACKEND_PASSWORD")
	hash, _ := bcrypt.Hash(payload.Password, s.Salt) //hash password
	payload.Password = hash
	return payload

	//if bcrypt.Match(payload.Password, hash) { //comprobacion de que matchean
	//fmt.Println("They match")
	//}

}

//executes a sign in query in sorare api and returns a type interface query result
func loadSignInQuery(s string, payload TUser) interface{} {

	client := graphql.NewClient("https://api.sorare.com/graphql")

	req := graphql.NewRequest(s)

	//sign, _ := json.Marshal(payload)

	req.Var("input", payload)

	_, csrf, _ := getSalt()
	req.Header.Set("x-csrf-token", csrf)

	ctx := context.Background()

	var respData SignIn
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)

	}

	return respData

}
