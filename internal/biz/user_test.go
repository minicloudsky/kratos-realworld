package biz

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "abc12345"

	encryptEdPasspord := hashPassword(password)
	spew.Dump(encryptEdPasspord)
	//panic(encryptEdPasspord)
}

func TestVerifyPassword(t *testing.T) {
	password := "abc12345"
	encryptEdPasspord := hashPassword(password)
	spew.Dump(encryptEdPasspord)
	//panic(encryptEdPasspord)
	verifyResult := verifyPassword(password, encryptEdPasspord)
	a := assert.New(t)
	//assert.Equal(t, true, verifyResult)
	//assert.Truef(t, verifyResult, "")
	a.Equal(true, verifyResult)
	a.True(verifyResult, "verifyResult: ", verifyResult)
}
