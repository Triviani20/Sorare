package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSignInQuery(t *testing.T) {

	assert := assert.New(t)

	//Test correct login
	payload := setPayload()
	resp, _, err := LoadSignInQuery(LOGIN_JWT, payload)
	fmt.Printf("%+v\n", resp)
	assert.Nil(err)
	assert.Equal("tim-appleton", resp.SignIn.CurrentUser.Slug)

	//Test incorrect login
	var ip TUser
	ip.Email = "asfddrfh"
	ip.Password = "incorrectpassword"
	resp, _, err = LoadSignInQuery(LOGIN_JWT, ip)
	//fmt.Printf("%+v\n", resp)
	assert.Nil(err)
	assert.Equal("", resp.SignIn.CurrentUser.Slug)
}
func TestGetListPendingOffer(t *testing.T) {

	assert := assert.New(t)

	//Test correct delete
	payload := setPayload()
	resp, _, err := LoadSignInQuery(LOGIN_JWT, payload)
	lpo := GetListPendingOffer(resp.SignIn.CurrentUser)
	fmt.Printf("%+v\n", lpo)
	assert.Nil(err)
}

func TestDeletePendingOffer(t *testing.T) {

	assert := assert.New(t)
	var cp TCancelPayload
	//Test correct delete
	payload := setPayload()
	resp, _, err := LoadSignInQuery(LOGIN_JWT, payload)
	lpo := GetListPendingOffer(resp.SignIn.CurrentUser)
	cOferr := DeletePendingOffer(CANCELOFFERPENDING_MUTATION, resp.SignIn.CurrentUser.JwtToken.Token, cp, lpo)
	fmt.Printf("%+v\n", cOferr)
	assert.Nil(err)

	//Test incorrect delete

}
