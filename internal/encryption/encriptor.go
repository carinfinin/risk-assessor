package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/carinfinin/risk-assessor/internal/model"
	"io"
)

type Encryptor struct {
	keyProvider KeyProvider
}

type KeyProvider interface {
	GetCurrentKeyID() string
	GetKey(keyID string) ([]byte, error)
}

func NewEncryptor(keyProvider KeyProvider) *Encryptor {
	return &Encryptor{
		keyProvider: keyProvider,
	}
}

// EncryptData шифрует структуру ClientData
func (e *Encryptor) EncryptData(data *model.ClientData) ([]byte, string, error) {
	// 1. Сериализуем данные в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, "", fmt.Errorf("marshal client data: %w", err)
	}

	// 2. Получаем текущий ключ
	keyID := e.keyProvider.GetCurrentKeyID()
	key, err := e.keyProvider.GetKey(keyID)
	if err != nil {
		return nil, "", fmt.Errorf("get encryption key: %w", err)
	}

	// 3. Шифруем данные
	ciphertext, err := e.encryptAESGCM(jsonData, key)
	if err != nil {
		return nil, "", fmt.Errorf("encrypt data: %w", err)
	}

	return ciphertext, keyID, nil
}

// DecryptData расшифровывает данные обратно в ClientData
func (e *Encryptor) DecryptData(ciphertext []byte, keyID string) (*model.ClientData, error) {
	// 1. Получаем ключ
	key, err := e.keyProvider.GetKey(keyID)
	if err != nil {
		return nil, fmt.Errorf("get decryption key: %w", err)
	}

	// 2. Расшифровываем данные
	plaintext, err := e.decryptAESGCM(ciphertext, key)
	if err != nil {
		return nil, fmt.Errorf("decrypt data: %w", err)
	}

	// 3. Десериализуем JSON
	var clientData model.ClientData
	if err := json.Unmarshal(plaintext, &clientData); err != nil {
		return nil, fmt.Errorf("unmarshal client data: %w", err)
	}

	return &clientData, nil
}

// encryptAESGCM шифрует данные с помощью AES-GCM
func (e *Encryptor) encryptAESGCM(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// decryptAESGCM расшифровывает данные
func (e *Encryptor) decryptAESGCM(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
