package app

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"strconv"
)

func (s *Server) HealthCheck(c *fiber.Ctx) error {
	return c.Status(200).SendString("OK")
}

func (s *Server) GetKeys(c *fiber.Ctx) error {
	// keyType := c.FormValue("type")
	keySize := c.FormValue("size")
	keySizeInt := 2048
	if keySize != "" {
		keySizeInt, _ = strconv.Atoi(keySize)

	}
	privateKey, err := rsa.GenerateKey(rand.Reader, keySizeInt)
	if err != nil {
		log.Error("Error while generate key")
		return c.Status(fiber.StatusInternalServerError).SendString("Error while generateKey")
	}
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA private key",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA public key",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	})

	return c.JSON(fiber.Map{"private key": string(privateKeyPEM),
		"public key": string(publicKeyPEM)})

}
