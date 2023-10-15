package peda

import (
	"fmt"
	"testing"

	"github.com/whatsauth/watoken"
)

func TestUpdateGetData(t *testing.T) {
	mconn := SetConnection("MONGOULBI", "petapedia")
	datagedung := GetAllBangunanLineString(mconn, "petapedia")
	fmt.Println(datagedung)
}

func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	GCFPostHandler("PRIVATEKEY", "MONGOULBI", "petapedia", "user")
}
