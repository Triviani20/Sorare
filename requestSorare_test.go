package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalQuery(t *testing.T) {

	assert := assert.New(t)

	_, err := MarshalQuery(P, VARSIGNIN, "")
	assert.Nil(err)
	//fmt.Println(string(b))
}

func TestLoginRequest(t *testing.T) {

	assert := assert.New(t)
	rsp, err := LoginRequest()
	assert.Nil(err)
	fmt.Println(rsp)
}

func TestCRREquest(t *testing.T) {

	assert := assert.New(t)
	_, err := CRREquest()
	assert.Nil(err)
	//fmt.Println(rsp)
}
