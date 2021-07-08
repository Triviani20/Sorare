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
	//fmt.Printf("COOKIE: %+v\n", c)
	s := readSalt(salt) //salt in struct

	payload.Email = os.Getenv("BACKEND_USER")
	payload.Password = os.Getenv("BACKEND_PASSWORD")
	hash, _ := bcrypt.Hash(payload.Password, s.Salt) //hash password
	payload.Password = hash
	payload.ClientMutationId = "Trivi"
	return payload

}

//executes a sign in query in sorare api and returns a type interface query result
func LoadSignInQuery(s string, payload TUser) (SignIn, string, error) {

	client := graphql.NewClient("https://api.sorare.com/graphql")

	req := graphql.NewRequest(s)

	req.Var("input", payload)

	_, csrf, _ := getSalt()

	req.Header.Set("X-CSRF-Token", csrf)
	req.Header.Set("Content-Type", "application/json")
	ctx := context.Background()

	var respData SignIn
	if err := client.Run(ctx, req, &respData); err != nil {
		return SignIn{}, "", err

	}
	return respData, csrf, nil

}

// devuelve una lista con todos los blockchainIds de todas las pending offers
// introducimos un tipo current user
func GetListPendingOffer(cU TCUser) (listPendingOffers []string) {

	pendOffersIds := cU.PendingDirectOffersSent.Nodes
	for _, n := range pendOffersIds {

		listPendingOffers = append(listPendingOffers, n.BlockchainId)
	}
	return listPendingOffers

}

func DeletePendingOffer(s string, jwt string, cancelPayload TCancelPayload, listPendingOffersToDelete []string) Cancel {

	client := graphql.NewClient("https://api.sorare.com/graphql")

	var respData Cancel

	for i := 0; i <= len(listPendingOffersToDelete); i++ {
		cancelPayload.BlockchainId = listPendingOffersToDelete[i]

		req := graphql.NewRequest(s)
		req.Var("input2", cancelPayload)

		//req.Header.Set("X-CSRF-Token", csrf)
		req.Header.Set("Content-Type", "application/json")

		bearer := "Bearer " + jwt
		fmt.Println(bearer)
		req.Header.Set("Authorization", bearer)
		req.Header.Set("JWT_AUD", AUD)
		ctx := context.Background()

		if err := client.Run(ctx, req, &respData); err != nil {
			log.Fatal(err)

		}
		fmt.Println("RESPDATA:", respData)
	}
	return respData
}
