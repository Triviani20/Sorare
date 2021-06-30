package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// make a get and take a salt
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
