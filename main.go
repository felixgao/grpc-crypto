package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	client "github.com/felix/grpc-crypto/client"
)

func main() {
	host := os.Getenv("IDPS_CLIENT_SERVICE_HOST")
	if len(host) == 0 {
		host = "localhost:11443"
	}
	conn, err := client.getIDPSConnection(host)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := client.NewCryptoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	fmt.Printf("Running Client\n")
	encryptedResponse, err := client.idpsEncrypt(ctx, client, "12345", "Felix Test")
	fmt.Printf("Encrypted Data:'" + encryptedResponse.GetEncryptedData() + "'\n")

	decryptedResponse, err := client.idpsDecrypt(ctx, client, "12345", "Felix Test-_-")
	fmt.Printf("Decrypted Data:'" + decryptedResponse.GetTextData() + "'\n")
}
