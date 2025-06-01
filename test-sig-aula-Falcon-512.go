package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"github.com/open-quantum-safe/liboqs-go/oqs"
)

func main() {
	// fmt.Println("liboqs version: "" + oqs.LiboqsVersion())
	// fmt.Println(oqs.EnabledSigs())

	// PASSO 1: Inicializar oqs.Signature
	signer := oqs.Signature{}
	err := signer.Init("Falcon-512", nil)
	if err != nil {
		log.Fatal("Erro ao inicializar o algoritmo: ", err)
	}

	// PASSO 2: Gerar chaves
	pubKey, err := signer.GenerateKeyPair()
	if err != nil {
		log.Fatal("Erro ao gerar chave: ", err)
	}
	// privateKey, err := signer.GetPrivateKey()
	// if err != nil {
	// 	log.Fatal("Erro ao obter chave privada: ", err)
	// }
	
	fmt.Println("Chave pública (Base64):")
	fmt.Println(base64.StdEncoding.EncodeToString(pubKey))
	// fmt.Println("Chave privada (Base64):")
	// fmt.Println(base64.StdEncoding.EncodeToString(privateKey))
	defer signer.Clean() // boa prática para liberar recursos
	
	// PASSO 3: Assinar uma mensagem
	message := []byte("mensagem de teste para assinatura post-quantum")
	signature, err := signer.Sign(message)
	if err != nil {
		log.Fatal("Erro ao assinar: ", err)
	}
	fmt.Println(string(message))
	fmt.Println("Assinatura (Base64):")
	fmt.Println(base64.StdEncoding.EncodeToString(signature))

	// PASSO 4: Verificar a assinatura
	verifier := oqs.Signature{}
	err = verifier.Init("Falcon-512", nil)
	if err != nil {
		log.Fatal("Erro ao inicializar o verificador: ", err)
	}
	valid, err := signer.Verify(message, signature, pubKey)
	if err != nil {
		log.Fatal("Erro na verificação: ", err)
	}
	if valid {
		fmt.Println("Assinatura verificada com sucesso!")
	} else {
		fmt.Println("Falha na verificação da assinatura.")
	}
	defer verifier.Clean() // boa prática para liberar recursos
}

