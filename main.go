package main

import (
	"cryptoForGo/crypt"
	"fmt"
)

func main() {
	key := []byte("123456781234567812345678")
	data := []byte("心有猛虎，细嗅蔷薇")

	cipherData, err := crypt.TripleDESEnCrypt(data, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(cipherData)
	origin, _ := crypt.TripleDESDeCrypt(cipherData, key)
	fmt.Println(string(origin))

}
