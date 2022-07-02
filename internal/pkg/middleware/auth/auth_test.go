package auth

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token := GenerateToken("hello", "minicloudsky")
	spew.Dump("token: ", token)
	fmt.Println("token: ", token)
	//panic(token)
}
