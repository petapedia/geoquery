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

func TestGeneratePasswordHash(t *testing.T) {
	password := "secret"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
}
