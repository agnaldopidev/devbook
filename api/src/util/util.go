package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

// Gerar chave key
func GeraKey() string {
	chave := make([]byte, 64)
	if _, err := rand.Read(chave); err != nil {
		log.Fatal(err)
		return ""
	}
	keyBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println("Chave gerada", keyBase64)
	return keyBase64
}
