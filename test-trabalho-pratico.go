package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/open-quantum-safe/liboqs-go/oqs"
)

func main() {
	// Ler a chave privada codificada em base64
	privKeyB64, err := ioutil.ReadFile("private_key.b64")
	if err != nil {
		log.Fatalf("Erro ao ler private_key.b64: %v", err)
	}
	// fmt.Println("Chave privada (Base64):")
	// fmt.Println(string(privKeyB64))

	privKey, err := base64.StdEncoding.DecodeString(string(privKeyB64))
	if err != nil {
		log.Fatalf("Erro ao decodificar chave privada: %v", err)
	}
	// fmt.Println("Chave privada (Base64):")
	// fmt.Println(base64.StdEncoding.EncodeToString(privKey))

	// Ler o ciphertext codificado em base64
	ctB64, err := ioutil.ReadFile("ciphertext.b64")
	if err != nil {
		log.Fatalf("Erro ao ler ciphertext.b64: %v", err)
	}
	// fmt.Println("Ciphertext (Base64):")
	// fmt.Println(string(ctB64))

	ciphertext, err := base64.StdEncoding.DecodeString(string(ctB64))
	if err != nil {
		log.Fatalf("Erro ao decodificar ciphertext: %v", err)
	}
	// fmt.Println("Ciphertext (Base64):")
	// fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))
	
	// Inicializar o KEM com Kyber768 e a chave privada
	kem := oqs.KeyEncapsulation{}
	if err := kem.Init("Kyber768", privKey); err != nil {
		log.Fatalf("Erro ao inicializar Kyber768: %v", err)
	}
	defer kem.Clean()

	// Decapsular (decifrar) o segredo
	sharedSecret, err := kem.DecapSecret(ciphertext)
	if err != nil {
		log.Fatalf("Erro ao decapsular o ciphertext: %v", err)
	}

	// Exibir segredo compartilhado (em base64)
	fmt.Println("Segredo compartilhado (Base64):")
	fmt.Println(base64.StdEncoding.EncodeToString(sharedSecret))
}
