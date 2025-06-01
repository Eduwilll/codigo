package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"github.com/open-quantum-safe/liboqs-go/oqs"
)

func main() {
	// fmt.Println(oqs.LiboqsVersion())
	// fmt.Println(oqs.EnabledSigs())

	// PASSO 1: Inicializar oqs.Signature
	signer := oqs.Signature{}
	err := signer.Init("Dilithium2", nil) // Substitua "Dilithium2" pelo algoritmo desejado e disponível
	if err != nil {
		log.Fatal("Erro ao inicializar o algoritmo: ", err)
	}
	defer signer.Clean() // boa prática para liberar recursos

	// PASSO 2: Gerar chaves
	pubKey, err := signer.GenerateKeyPair()
	if err != nil {
		log.Fatal("Erro ao gerar chave: ", err)
	}
	secretKey := signer.ExportSecretKey()
	

	fmt.Println("Chave privada (Base64):")
	fmt.Println(base64.StdEncoding.EncodeToString(secretKey))
	// fmt.Println("Chave privada (Hex):")
	// fmt.Println(hex.EncodeToString(ExportSecretKey))
	// fmt.Println("Chave pública (Hex):")
	// fmt.Println(hex.EncodeToString(pubKey))
	// fmt.Println("Chave pública (Base64):")
	// fmt.Println(base64.StdEncoding.EncodeToString(pubKey))
	fmt.Println("Chave pública (Base64):")
	fmt.Println(base64.StdEncoding.EncodeToString(pubKey))

	// PASSO 3: Assinar uma mensagem
	message := []byte("mensagem de teste para assinatura post-quantum")
	signature, err := signer.Sign(message)
	if err != nil {
		log.Fatal("Erro ao assinar: ", err)
	}
	fmt.Println("Assinatura (Base64):")
	fmt.Println(base64.StdEncoding.EncodeToString(signature))

	// PASSO 4: Verificar a assinatura
	valid, err := signer.Verify(message, signature, pubKey)
	if err != nil {
		log.Fatal("Erro na verificação: ", err)
	}
	if valid {
		fmt.Println("Assinatura verificada com sucesso!")
	} else {
		fmt.Println("Falha na verificação da assinatura.")
	}
}

