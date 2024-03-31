package app

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type EncryptRequestDecode struct {
	Data       string `json:"data"`
	PrivateKey string `json:"private_key"`
}

func (s *Server) DecryptData(c *fiber.Ctx) error {
	// Прочитать тело запроса в структуру EncryptRequest
	var req EncryptRequestDecode
	if err := c.BodyParser(&req); err != nil {
		log.Error("Ошибка чтения запроса:", err)
		return c.Status(fiber.StatusBadRequest).SendString("Ошибка чтения запроса")
	}

	data, err := hex.DecodeString(req.Data)
	// Получение данных для расшифровки
	requestData := []byte(data)

	// Получение закрытого ключа и декодирование его из PEM
	privateKeyPEM := []byte(req.PrivateKey)
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		log.Error("Ошибка декодирования PEM блока")
		return c.Status(fiber.StatusBadRequest).SendString("Ошибка декодирования PEM блока")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Error("Ошибка парсинга закрытого ключа:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка парсинга закрытого ключа")
	}

	label := []byte("")
	// Расшифрование данных с использованием закрытого ключа
	decryptedData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, requestData, label)
	if err != nil {
		log.Error("Ошибка расшифровки данных:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка расшифровки данных")
	}

	return c.JSON(fiber.Map{"decrypted_data": string(decryptedData)})
}
