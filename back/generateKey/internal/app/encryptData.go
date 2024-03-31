package app

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type EncryptRequest struct {
	Data      string `json:"data"`
	PublicKey string `json:"public_key"`
}

func (s *Server) EncryptData(c *fiber.Ctx) error {
	// Прочитать тело запроса в структуру EncryptRequest
	var req EncryptRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error("Ошибка чтения запроса:", err)
		return c.Status(fiber.StatusBadRequest).SendString("Ошибка чтения запроса")
	}

	// Получение данных для шифрования
	requestData := []byte(req.Data)

	// Получение публичного ключа и декодирование его из PEM
	publicKeyPEM := []byte(req.PublicKey)
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil {
		log.Error("Ошибка декодирования PEM блока")
		return c.Status(fiber.StatusBadRequest).SendString("Ошибка декодирования PEM блока")
	}
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.Error("Ошибка парсинга публичного ключа:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка парсинга публичного ключа")
	}

	label := []byte("")
	// Шифрование данных с использованием открытого ключа
	encryptedData, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, requestData, label)
	if err != nil {
		log.Error("Ошибка шифрования данных:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка шифрования данных")
	}

	return c.JSON(fiber.Map{"encrypted_data": fmt.Sprintf("%x", encryptedData)})
}
